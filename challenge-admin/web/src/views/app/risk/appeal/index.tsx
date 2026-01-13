import { CloudDownloadOutlined, CheckOutlined, StopOutlined } from "@ant-design/icons";
import { ProColumns, ProFormInstance, ProTable } from "@ant-design/pro-components";
import { Button, Modal, Space } from "antd";
import React, { useRef } from "react";
import HocAuth from "@/components/HocAuth";
import LoadingButton from "@/components/LoadingButton";
import { pagination } from "@/config/proTable";
import { formatDataForProTable, saveExcelBlob } from "@/utils";
import { message, modal } from "@/hooks/useMessage";
import { RiskAppealModel, getRiskAppealPageApi, exportRiskAppealApi, reviewRiskAppealApi } from "@/api/app/risk/appeal";
import type { ActionType } from "@ant-design/pro-components";

const RiskAppealPage: React.FC = () => {
	const actionRef = useRef<ActionType>();
	const formRef = useRef<ProFormInstance>();

	const handleExport = async (done: () => void) => {
		try {
			saveExcelBlob("risk_appeal", await exportRiskAppealApi(formRef.current?.getFieldsValue()));
		} catch (err) {
			message.error("导出失败");
		} finally {
			done();
		}
	};

	const handleReview = (id: number, status: number) => {
		Modal.confirm({
			title: status === 2 ? "确认通过申诉？" : "确认拒绝申诉？",
			onOk: async () => {
				await reviewRiskAppealApi(id, { status, actionResult: status === 2 ? 1 : 2, reviewRemark: "" });
				message.success("操作成功");
				actionRef.current?.reload();
			}
		});
	};

	const columns: ProColumns<RiskAppealModel>[] = [
		{ title: "申诉ID", dataIndex: "id", width: 100 },
		{ title: "用户ID", dataIndex: "userId", width: 120 },
		{
			title: "状态",
			dataIndex: "status",
			valueType: "select",
			valueEnum: {
				1: { text: "待处理" },
				2: { text: "通过" },
				3: { text: "拒绝" }
			},
			width: 120
		},
		{ title: "设备指纹", dataIndex: "deviceFp", width: 200 },
		{ title: "申诉理由", dataIndex: "appealReason", width: 220, hideInSearch: true, ellipsis: true },
		{ title: "申诉时间", dataIndex: "createdAt", valueType: "dateTime", width: 180, hideInSearch: true },
		{ title: "审核时间", dataIndex: "reviewedAt", valueType: "dateTime", width: 180, hideInSearch: true },
		{
			title: "操作",
			valueType: "option",
			width: 220,
			fixed: "right",
			render: (_, record) => (
				<Space>
					<HocAuth permission={["app:risk:appeal:review"]}>
						<Button type="primary" size="small" icon={<CheckOutlined />} disabled={record.status !== 1} onClick={() => handleReview(record.id!, 2)}>
							通过
						</Button>
						<Button danger size="small" icon={<StopOutlined />} disabled={record.status !== 1} onClick={() => handleReview(record.id!, 3)}>
							拒绝
						</Button>
					</HocAuth>
				</Space>
			)
		}
	];

	return (
		<ProTable<RiskAppealModel>
			rowKey="id"
			columns={columns}
			actionRef={actionRef}
			formRef={formRef}
			pagination={pagination}
			scroll={{ x: 1400 }}
			request={async params => {
				const { data } = await getRiskAppealPageApi(params);
				return formatDataForProTable<RiskAppealModel>(data);
			}}
			toolBarRender={() => [
				<HocAuth key="export" permission={["app:risk:appeal:export"]}>
					<LoadingButton type="primary" key="exportBtn" icon={<CloudDownloadOutlined />} onClick={done => handleExport(done)}>
						导出
					</LoadingButton>
				</HocAuth>
			]}
			search={{ labelWidth: "auto" }}
		/>
	);
};

export default RiskAppealPage;
