import { CloudDownloadOutlined } from "@ant-design/icons";
import { ProColumns, ProFormInstance, ProTable } from "@ant-design/pro-components";
import React, { useRef } from "react";
import HocAuth from "@/components/HocAuth";
import LoadingButton from "@/components/LoadingButton";
import { pagination } from "@/config/proTable";
import { formatDataForProTable, saveExcelBlob } from "@/utils";
import { message } from "@/hooks/useMessage";
import { RiskDeviceModel, getRiskDevicePageApi, exportRiskDeviceApi } from "@/api/app/risk/device";
import type { ActionType } from "@ant-design/pro-components";

const RiskDevicePage: React.FC = () => {
	const actionRef = useRef<ActionType>();
	const formRef = useRef<ProFormInstance>();

	const handleExport = async (done: () => void) => {
		try {
			saveExcelBlob("risk_device", await exportRiskDeviceApi(formRef.current?.getFieldsValue()));
		} catch (err) {
			message.error("导出失败");
		} finally {
			done();
		}
	};

	const columns: ProColumns<RiskDeviceModel>[] = [
		{ title: "ID", dataIndex: "id", width: 80 },
		{ title: "设备指纹", dataIndex: "deviceFp", width: 260 },
		{ title: "用户ID", dataIndex: "userId", width: 120 },
		{ title: "记录时间", dataIndex: "createdAt", valueType: "dateTime", width: 180, hideInSearch: true }
	];

	return (
		<ProTable<RiskDeviceModel>
			rowKey="id"
			columns={columns}
			actionRef={actionRef}
			formRef={formRef}
			pagination={pagination}
			scroll={{ x: 1200 }}
			request={async params => {
				const { data } = await getRiskDevicePageApi(params);
				return formatDataForProTable<RiskDeviceModel>(data);
			}}
			toolBarRender={() => [
				<HocAuth key="export" permission={["app:risk:device:export"]}>
					<LoadingButton type="primary" key="exportBtn" icon={<CloudDownloadOutlined />} onClick={done => handleExport(done)}>
						导出
					</LoadingButton>
				</HocAuth>
			]}
			search={{ labelWidth: "auto" }}
		/>
	);
};

export default RiskDevicePage;
