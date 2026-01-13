import { CloudDownloadOutlined } from "@ant-design/icons";
import { ProColumns, ProFormInstance, ProTable } from "@ant-design/pro-components";
import React, { useRef } from "react";
import HocAuth from "@/components/HocAuth";
import LoadingButton from "@/components/LoadingButton";
import { pagination } from "@/config/proTable";
import { formatDataForProTable, saveExcelBlob } from "@/utils";
import { message } from "@/hooks/useMessage";
import { RiskUserModel, getRiskUserPageApi, exportRiskUserApi } from "@/api/app/risk/user";
import type { ActionType } from "@ant-design/pro-components";

const RiskUserPage: React.FC = () => {
	const actionRef = useRef<ActionType>();
	const formRef = useRef<ProFormInstance>();

	const handleExport = async (done: () => void) => {
		try {
			saveExcelBlob("risk_user", await exportRiskUserApi(formRef.current?.getFieldsValue()));
		} catch (err) {
			message.error("导出失败");
		} finally {
			done();
		}
	};

	const columns: ProColumns<RiskUserModel>[] = [
		{ title: "用户ID", dataIndex: "userId", width: 160 },
		{
			title: "风险等级",
			dataIndex: "riskLevel",
			valueType: "select",
			valueEnum: {
				0: { text: "正常" },
				1: { text: "观察" },
				2: { text: "限制" },
				3: { text: "封禁" }
			},
			width: 120
		},
		{ title: "风险分", dataIndex: "riskScore", width: 120, hideInSearch: true },
		{ title: "风险原因", dataIndex: "reason", width: 220, hideInSearch: true, ellipsis: true },
		{ title: "更新时间", dataIndex: "updatedAt", valueType: "dateTime", width: 180, hideInSearch: true }
	];

	return (
		<ProTable<RiskUserModel>
			rowKey="userId"
			columns={columns}
			actionRef={actionRef}
			formRef={formRef}
			pagination={pagination}
			scroll={{ x: 1200 }}
			request={async params => {
				const { data } = await getRiskUserPageApi(params);
				return formatDataForProTable<RiskUserModel>(data);
			}}
			toolBarRender={() => [
				<HocAuth key="export" permission={["app:risk:user:export"]}>
					<LoadingButton type="primary" key="exportBtn" icon={<CloudDownloadOutlined />} onClick={done => handleExport(done)}>
						导出
					</LoadingButton>
				</HocAuth>
			]}
			search={{ labelWidth: "auto" }}
		/>
	);
};

export default RiskUserPage;
