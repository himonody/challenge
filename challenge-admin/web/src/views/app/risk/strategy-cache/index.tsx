import { CloudDownloadOutlined } from "@ant-design/icons";
import { ProColumns, ProFormInstance, ProTable } from "@ant-design/pro-components";
import React, { useRef } from "react";
import HocAuth from "@/components/HocAuth";
import LoadingButton from "@/components/LoadingButton";
import { pagination } from "@/config/proTable";
import { formatDataForProTable, saveExcelBlob } from "@/utils";
import { message } from "@/hooks/useMessage";
import { RiskStrategyCacheModel, getRiskStrategyCachePageApi, exportRiskStrategyCacheApi } from "@/api/app/risk/strategyCache";
import type { ActionType } from "@ant-design/pro-components";

const RiskStrategyCachePage: React.FC = () => {
	const actionRef = useRef<ActionType>();
	const formRef = useRef<ProFormInstance>();

	const handleExport = async (done: () => void) => {
		try {
			saveExcelBlob("risk_strategy_cache", await exportRiskStrategyCacheApi(formRef.current?.getFieldsValue()));
		} catch (err) {
			message.error("导出失败");
		} finally {
			done();
		}
	};

	const columns: ProColumns<RiskStrategyCacheModel>[] = [
		{ title: "场景", dataIndex: "scene", width: 160 },
		{ title: "统计维度", dataIndex: "identityType", width: 140 },
		{ title: "规则编码", dataIndex: "ruleCode", width: 200 },
		{ title: "窗口秒", dataIndex: "windowSeconds", width: 120, hideInSearch: true },
		{ title: "阈值次数", dataIndex: "threshold", width: 120, hideInSearch: true },
		{ title: "动作编码", dataIndex: "action", width: 160, hideInSearch: true },
		{ title: "动作值", dataIndex: "actionValue", width: 120, hideInSearch: true }
	];

	return (
		<ProTable<RiskStrategyCacheModel>
			rowKey="ruleCode"
			columns={columns}
			actionRef={actionRef}
			formRef={formRef}
			pagination={pagination}
			scroll={{ x: 1400 }}
			request={async params => {
				const { data } = await getRiskStrategyCachePageApi(params);
				return formatDataForProTable<RiskStrategyCacheModel>(data);
			}}
			toolBarRender={() => [
				<HocAuth key="export" permission={["app:risk:strategyCache:export"]}>
					<LoadingButton type="primary" key="exportBtn" icon={<CloudDownloadOutlined />} onClick={done => handleExport(done)}>
						导出
					</LoadingButton>
				</HocAuth>
			]}
			search={{ labelWidth: "auto" }}
		/>
	);
};

export default RiskStrategyCachePage;
