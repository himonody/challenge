import { ExclamationCircleOutlined, PlusCircleOutlined, CloudDownloadOutlined, EditOutlined } from "@ant-design/icons";
import { ProColumns, ProFormInstance, ProTable } from "@ant-design/pro-components";
import { Modal, Space } from "antd";
import React, { useRef } from "react";
import HocAuth from "@/components/HocAuth";
import LoadingButton from "@/components/LoadingButton";
import { pagination } from "@/config/proTable";
import { formatDataForProTable, saveExcelBlob } from "@/utils";
import { message, modal } from "@/hooks/useMessage";
import { RiskBlacklistModel, getRiskBlacklistPageApi, exportRiskBlacklistApi } from "@/api/app/risk/blacklist";
import type { ActionType } from "@ant-design/pro-components";
import FormModal from "./components/FormModal";
import type { FormModalRef } from "./components/FormModal";

const RiskBlacklistPage: React.FC = () => {
	const actionRef = useRef<ActionType>();
	const formRef = useRef<ProFormInstance>();
	const formModalRef = useRef<FormModalRef>(null);

	const handleExport = (done: () => void) => {
		modal.confirm({
			title: "导出提示",
			icon: <ExclamationCircleOutlined />,
			content: "确认导出当前筛选数据？",
			onCancel: () => done(),
			onOk: async () => {
				try {
					saveExcelBlob("risk_blacklist", await exportRiskBlacklistApi(formRef.current?.getFieldsValue()));
				} catch (err) {
					message.error("导出失败");
				} finally {
					done();
				}
			}
		});
	};

	const handleShowCreate = () => formModalRef.current?.showAddFormModal();
	const handleShowEdit = (record: RiskBlacklistModel) => {
		if (record.id === undefined) return;
		formModalRef.current?.showEditFormModal(record.id, record);
	};

	const columns: ProColumns<RiskBlacklistModel>[] = [
		{ title: "ID", dataIndex: "id", width: 80 },
		{
			title: "类型",
			dataIndex: "type",
			valueType: "select",
			valueEnum: {
				ip: { text: "IP" },
				device: { text: "设备" },
				country: { text: "国家" },
				mobile: { text: "手机" },
				email: { text: "邮箱" }
			},
			width: 120
		},
		{ title: "命中值", dataIndex: "value", width: 200 },
		{ title: "风险等级", dataIndex: "riskLevel", width: 100 },
		{ title: "封禁原因", dataIndex: "reason", width: 200, ellipsis: true, hideInSearch: true },
		{
			title: "状态",
			dataIndex: "status",
			valueType: "select",
			valueEnum: {
				"1": { text: "生效" },
				"2": { text: "失效" }
			},
			width: 100
		},
		{ title: "创建时间", dataIndex: "createdAt", valueType: "dateTime", width: 180, hideInSearch: true },
		{
			title: "操作",
			valueType: "option",
			width: 160,
			fixed: "right",
			render: (_, record) => (
				<Space>
					<HocAuth permission={["app:risk:blacklist:edit"]}>
						<LoadingButton key="edit" type="link" size="small" icon={<EditOutlined />} onClick={() => handleShowEdit(record)}>
							编辑
						</LoadingButton>
					</HocAuth>
				</Space>
			)
		}
	];

	const toolBarRender = () => [
		<HocAuth key="add" permission={["app:risk:blacklist:add"]}>
			<LoadingButton type="primary" key="addBtn" icon={<PlusCircleOutlined />} onClick={() => handleShowCreate()}>
				新增
			</LoadingButton>
		</HocAuth>,
		<HocAuth key="export" permission={["app:risk:blacklist:export"]}>
			<LoadingButton type="primary" key="exportBtn" icon={<CloudDownloadOutlined />} onClick={done => handleExport(done)}>
				导出
			</LoadingButton>
		</HocAuth>
	];

	return (
		<>
			<ProTable<RiskBlacklistModel>
				rowKey="id"
				columns={columns}
				actionRef={actionRef}
				formRef={formRef}
				pagination={pagination}
				scroll={{ x: 1200 }}
				request={async params => {
					const { data } = await getRiskBlacklistPageApi(params);
					return formatDataForProTable<RiskBlacklistModel>(data);
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

export default RiskBlacklistPage;
