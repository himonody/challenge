import { CloudDownloadOutlined } from "@ant-design/icons";
import { ProColumns, ProFormInstance, ProTable } from "@ant-design/pro-components";
import React, { useRef } from "react";
import HocAuth from "@/components/HocAuth";
import LoadingButton from "@/components/LoadingButton";
import { pagination } from "@/config/proTable";
import { formatDataForProTable, saveExcelBlob } from "@/utils";
import { message } from "@/hooks/useMessage";
import { RiskEventModel, getRiskEventPageApi, exportRiskEventApi } from "@/api/app/risk/event";
import type { ActionType } from "@ant-design/pro-components";

const RiskEventPage: React.FC = () => {
	const actionRef = useRef<ActionType>();
	const formRef = useRef<ProFormInstance>();

	const handleExport = async (done: () => void) => {
		try {
			saveExcelBlob("risk_event", await exportRiskEventApi(formRef.current?.getFieldsValue()));
		} catch (err) {
			message.error("导出失败");
		} finally {
			done();
		}
	};

	const columns: ProColumns<RiskEventModel>[] = [
		{ title: "事件ID", dataIndex: "id", width: 120 },
		{ title: "用户ID", dataIndex: "userId", width: 120 },
		{ title: "事件类型", dataIndex: "eventType", width: 100 },
		{ title: "事件详情", dataIndex: "detail", width: 200, hideInSearch: true },
		{ title: "风险分", dataIndex: "score", width: 120, hideInSearch: true },
		{ title: "发生时间", dataIndex: "createdAt", valueType: "dateTime", width: 180, hideInSearch: true }
	];

	return (
		<ProTable<RiskEventModel>
			rowKey="id"
			columns={columns}
			actionRef={actionRef}
			formRef={formRef}
			pagination={pagination}
			scroll={{ x: 1300 }}
			request={async params => {
				const { data } = await getRiskEventPageApi(params);
				return formatDataForProTable<RiskEventModel>(data);
			}}
			toolBarRender={() => [
				<HocAuth key="export" permission={["app:risk:event:export"]}>
					<LoadingButton type="primary" key="exportBtn" icon={<CloudDownloadOutlined />} onClick={done => handleExport(done)}>
						导出
					</LoadingButton>
				</HocAuth>
			]}
			search={{ labelWidth: "auto" }}
		/>
	);
};

export default RiskEventPage;
