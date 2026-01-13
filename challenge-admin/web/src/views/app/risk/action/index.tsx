import { ExclamationCircleOutlined, PlusCircleOutlined, CloudDownloadOutlined, EditOutlined, DeleteOutlined } from "@ant-design/icons";
import { ProColumns, ProFormInstance, ProTable } from "@ant-design/pro-components";
import { Modal, Space } from "antd";
import React, { useRef } from "react";
import HocAuth from "@/components/HocAuth";
import LoadingButton from "@/components/LoadingButton";
import { pagination } from "@/config/proTable";
import { formatDataForProTable, saveExcelBlob } from "@/utils";
import { message, modal } from "@/hooks/useMessage";
import { RiskActionModel, getRiskActionPageApi, exportRiskActionApi, delRiskActionApi } from "@/api/app/risk/action";
import type { ActionType } from "@ant-design/pro-components";
import FormModal from "./components/FormModal";
import type { FormModalRef } from "./components/FormModal";

const RiskActionPage: React.FC = () => {
	const actionRef = useRef<ActionType>();
	const formRef = useRef<ProFormInstance>();
	const formModalRef = useRef<FormModalRef>(null);

	const handleDelete = (code: string) => {
		Modal.confirm({
			title: "确认删除该动作？",
			icon: <ExclamationCircleOutlined />,
			onOk: async () => {
				await delRiskActionApi(code);
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
					saveExcelBlob("risk_action", await exportRiskActionApi(formRef.current?.getFieldsValue()));
				} catch (err) {
					message.error("导出失败");
				} finally {
					done();
				}
			}
		});
	};

	const handleShowCreate = () => formModalRef.current?.showAddFormModal();
	const handleShowEdit = (record: RiskActionModel) => {
		if (!record.code) return;
		formModalRef.current?.showEditFormModal(record.code, record);
	};

	const columns: ProColumns<RiskActionModel>[] = [
		{ title: "动作编码", dataIndex: "code", width: 160 },
		{ title: "动作类型", dataIndex: "type", width: 120 },
		{ title: "默认动作值", dataIndex: "defaultValue", width: 120 },
		{ title: "说明", dataIndex: "remark", width: 200, ellipsis: true },
		{ title: "创建时间", dataIndex: "createdAt", valueType: "dateTime", width: 180, hideInSearch: true },
		{
			title: "操作",
			valueType: "option",
			width: 200,
			fixed: "right",
			render: (_, record) => (
				<Space>
					<HocAuth permission={["app:risk:action:edit"]}>
						<LoadingButton key="edit" type="link" size="small" icon={<EditOutlined />} onClick={() => handleShowEdit(record)}>
							编辑
						</LoadingButton>
					</HocAuth>
					<HocAuth permission={["app:risk:action:delete"]}>
						<LoadingButton key="del" danger type="link" size="small" icon={<DeleteOutlined />} onClick={() => handleDelete(record.code!)}>
							删除
						</LoadingButton>
					</HocAuth>
				</Space>
			)
		}
	];

	const toolBarRender = () => [
		<HocAuth key="add" permission={["app:risk:action:add"]}>
			<LoadingButton type="primary" key="addBtn" icon={<PlusCircleOutlined />} onClick={() => handleShowCreate()}>
				新增
			</LoadingButton>
		</HocAuth>,
		<HocAuth key="export" permission={["app:risk:action:export"]}>
			<LoadingButton type="primary" key="exportBtn" icon={<CloudDownloadOutlined />} onClick={done => handleExport(done)}>
				导出
			</LoadingButton>
		</HocAuth>
	];

	return (
		<>
			<ProTable<RiskActionModel>
				rowKey="code"
				columns={columns}
				actionRef={actionRef}
				formRef={formRef}
				pagination={pagination}
				scroll={{ x: 1200 }}
				request={async params => {
					const { data } = await getRiskActionPageApi(params);
					return formatDataForProTable<RiskActionModel>(data);
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

export default RiskActionPage;
