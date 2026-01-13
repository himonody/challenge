import { CloudDownloadOutlined } from "@ant-design/icons";
import { ProColumns, ProFormInstance, ProTable } from "@ant-design/pro-components";
import React, { useRef } from "react";
import HocAuth from "@/components/HocAuth";
import LoadingButton from "@/components/LoadingButton";
import { pagination } from "@/config/proTable";
import { formatDataForProTable, saveExcelBlob } from "@/utils";
import { message } from "@/hooks/useMessage";
import { RiskRateLimitModel, getRiskRateLimitPageApi, exportRiskRateLimitApi } from "@/api/app/risk/rateLimit";
import type { ActionType } from "@ant-design/pro-components";

const RateLimitPage: React.FC = () => {
	const actionRef = useRef<ActionType>();
	const formRef = useRef<ProFormInstance>();

	const handleExport = async (done: () => void) => {
		try {
			saveExcelBlob("risk_rate_limit", await exportRiskRateLimitApi(formRef.current?.getFieldsValue()));
		} catch (err) {
			message.error("导出失败");
		} finally {
			done();
		}
	};

	const columns: ProColumns<RiskRateLimitModel>[] = [
		{ title: "ID", dataIndex: "id", width: 80 },
		{ title: "场景", dataIndex: "scene", width: 140 },
		{ title: "标识类型", dataIndex: "identityType", width: 120 },
		{ title: "标识值", dataIndex: "identityValue", width: 200 },
		{ title: "次数", dataIndex: "count", width: 100, hideInSearch: true },
		{ title: "窗口开始", dataIndex: "windowStart", valueType: "dateTime", width: 180 },
		{ title: "窗口结束", dataIndex: "windowEnd", valueType: "dateTime", width: 180 },
		{
			title: "是否拦截",
			dataIndex: "blocked",
			valueType: "select",
			valueEnum: {
				"1": { text: "是" },
				"0": { text: "否" }
			},
			width: 100
		},
		{ title: "记录时间", dataIndex: "createdAt", valueType: "dateTime", width: 180, hideInSearch: true }
	];

	const toolBarRender = () => [
		<HocAuth key="export" permission={["app:risk:rateLimit:export"]}>
			<LoadingButton type="primary" key="exportBtn" icon={<CloudDownloadOutlined />} onClick={done => handleExport(done)}>
				导出
			</LoadingButton>
		</HocAuth>
	];

	return (
		<ProTable<RiskRateLimitModel>
			rowKey="id"
			columns={columns}
			actionRef={actionRef}
			formRef={formRef}
			pagination={pagination}
			scroll={{ x: 1400 }}
			request={async params => {
				const { data } = await getRiskRateLimitPageApi(params);
				return formatDataForProTable<RiskRateLimitModel>(data);
			}}
			toolBarRender={toolBarRender}
			search={{ labelWidth: "auto" }}
		/>
	);
};

export default RateLimitPage;
