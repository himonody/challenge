import { ExclamationCircleOutlined, PlusCircleOutlined, CloudDownloadOutlined, EditOutlined, DeleteOutlined } from "@ant-design/icons";
import { ProColumns, ProFormInstance, ProTable } from "@ant-design/pro-components";
import { Modal, Space } from "antd";
import React, { useRef } from "react";
import HocAuth from "@/components/HocAuth";
import LoadingButton from "@/components/LoadingButton";
import { pagination } from "@/config/proTable";
import { formatDataForProTable, saveExcelBlob } from "@/utils";
import { message, modal } from "@/hooks/useMessage";
import { RiskStrategyModel, getRiskStrategyPageApi, exportRiskStrategyApi, delRiskStrategyApi } from "@/api/app/risk/strategy";
import type { ActionType } from "@ant-design/pro-components";
import FormModal from "./components/FormModal";
import type { FormModalRef } from "./components/FormModal";

const RiskStrategyPage: React.FC = () => {
	const actionRef = useRef<ActionType>();
	const formRef = useRef<ProFormInstance>();
	const formModalRef = useRef<FormModalRef>(null);

	const handleDelete = (id: number) => {
		Modal.confirm({
			title: "确认删除该策略？",
			icon: <ExclamationCircleOutlined />,
			onOk: async () => {
				await delRiskStrategyApi(id);
				message.success("删除成功");
				actionRef.current?.reload();
			}
		});
	};

	const handleExport = (done: () => void) => {
		modal.confirm({
			title: "导出提示",
			icon: <ExclamationCircleOutlined />,
			content: "确认导出当前筛选数据？",
			onCancel: () => done(),
			onOk: async () => {
				try {
					saveExcelBlob("risk_strategy", await exportRiskStrategyApi(formRef.current?.getFieldsValue()));
				} catch (err) {
					message.error("导出失败");
				} finally {
					done();
				}
			}
		});
	};

	const handleShowCreate = () => formModalRef.current?.showAddFormModal();
	const handleShowEdit = (record: RiskStrategyModel) => {
		if (record.id === undefined) return;
		formModalRef.current?.showEditFormModal(record.id, record);
	};

	const columns: ProColumns<RiskStrategyModel>[] = [
		{ title: "ID", dataIndex: "id", width: 80 },
		{ title: "场景", dataIndex: "scene", width: 120 },
		{ title: "规则编码", dataIndex: "ruleCode", width: 140 },
		{ title: "统计维度", dataIndex: "identityType", width: 120 },
		{ title: "窗口秒", dataIndex: "windowSeconds", width: 100 },
		{ title: "阈值次数", dataIndex: "threshold", width: 100 },
		{ title: "动作编码", dataIndex: "action", width: 140 },
		{ title: "动作值", dataIndex: "actionValue", width: 100 },
		{
			title: "状态",
			dataIndex: "status",
			valueType: "select",
			valueEnum: {
				1: { text: "启用" },
				0: { text: "停用" }
			},
			width: 100
		},
		{ title: "优先级", dataIndex: "priority", width: 100 },
		{ title: "说明", dataIndex: "remark", width: 200, ellipsis: true, hideInSearch: true },
		{ title: "更新时间", dataIndex: "updatedAt", valueType: "dateTime", width: 180, hideInSearch: true },
		{
			title: "操作",
			valueType: "option",
			width: 200,
			fixed: "right",
			render: (_, record) => (
				<Space>
					<HocAuth permission={["app:risk:strategy:edit"]}>
						<LoadingButton key="edit" type="link" size="small" icon={<EditOutlined />} onClick={() => handleShowEdit(record)}>
							编辑
						</LoadingButton>
					</HocAuth>
					<HocAuth permission={["app:risk:strategy:delete"]}>
						<LoadingButton key="del" danger type="link" size="small" icon={<DeleteOutlined />} onClick={() => handleDelete(record.id!)}>
							删除
						</LoadingButton>
					</HocAuth>
				</Space>
			)
		}
	];

	const toolBarRender = () => [
		<HocAuth key="add" permission={["app:risk:strategy:add"]}>
			<LoadingButton type="primary" key="addBtn" icon={<PlusCircleOutlined />} onClick={() => handleShowCreate()}>
				新增
			</LoadingButton>
		</HocAuth>,
		<HocAuth key="export" permission={["app:risk:strategy:export"]}>
			<LoadingButton type="primary" key="exportBtn" icon={<CloudDownloadOutlined />} onClick={done => handleExport(done)}>
				导出
			</LoadingButton>
		</HocAuth>
	];

	return (
		<>
			<ProTable<RiskStrategyModel>
				rowKey="id"
				columns={columns}
				actionRef={actionRef}
				formRef={formRef}
				pagination={pagination}
				scroll={{ x: 1400 }}
				request={async params => {
					const { data } = await getRiskStrategyPageApi(params);
					return formatDataForProTable<RiskStrategyModel>(data);
				}}
				toolBarRender={toolBarRender}
				search={{ labelWidth: "auto" }}
			/>
			<FormModal
				ref={formModalRef}
				onSuccess={() => {
					actionRef.current?.reload();
				}}
			/>
		</>
	);
};

export default RiskStrategyPage;
