import { ExclamationCircleOutlined, PlusCircleOutlined, CloudDownloadOutlined, EditOutlined, DeleteOutlined } from "@ant-design/icons";
import { ProColumns, ProFormInstance, ProTable } from "@ant-design/pro-components";
import { Modal, Space } from "antd";
import React, { useRef } from "react";
import HocAuth from "@/components/HocAuth";
import LoadingButton from "@/components/LoadingButton";
import { pagination } from "@/config/proTable";
import { formatDataForProTable, saveExcelBlob } from "@/utils";
import { message, modal } from "@/hooks/useMessage";
import { ChallengeDailyStatModel, getDailyStatPageApi, exportDailyStatApi, delDailyStatApi } from "@/api/app/challenge/dailyStat";
import type { ActionType } from "@ant-design/pro-components";
import FormModal from "./components/FormModal";
import type { FormModalRef } from "./components/FormModal";

const DailyStatPage: React.FC = () => {
	const actionRef = useRef<ActionType>();
	const formRef = useRef<ProFormInstance>();
	const formModalRef = useRef<FormModalRef>(null);

	const handleDelete = (statDate: string) => {
		Modal.confirm({
			title: "确认删除该每日统计？",
			icon: <ExclamationCircleOutlined />,
			onOk: async () => {
				await delDailyStatApi(statDate);
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
					saveExcelBlob("challenge_daily_stat", await exportDailyStatApi(formRef.current?.getFieldsValue()));
				} catch (err) {
					message.error("导出失败");
				} finally {
					done();
				}
			}
		});
	};

	const handleShowCreate = () => formModalRef.current?.showAddFormModal();
	const handleShowEdit = (record: ChallengeDailyStatModel) => {
		if (!record.statDate) return;
		formModalRef.current?.showEditFormModal(record.statDate, record);
	};

	const columns: ProColumns<ChallengeDailyStatModel>[] = [
		{ title: "统计日期", dataIndex: "statDate", valueType: "date", width: 120 },
		{ title: "参与人数", dataIndex: "joinUserCnt", width: 100 },
		{ title: "成功人数", dataIndex: "successUserCnt", width: 100 },
		{ title: "失败人数", dataIndex: "failUserCnt", width: 100 },
		{ title: "参与金额", dataIndex: "joinAmount", width: 130, hideInSearch: true },
		{ title: "成功金额", dataIndex: "successAmount", width: 130, hideInSearch: true },
		{ title: "失败金额", dataIndex: "failAmount", width: 130, hideInSearch: true },
		{ title: "平台奖励", dataIndex: "platformBonus", width: 130, hideInSearch: true },
		{ title: "奖池金额", dataIndex: "poolAmount", width: 130, hideInSearch: true },
		{ title: "创建时间", dataIndex: "createdAt", valueType: "dateTime", width: 180, hideInSearch: true },
		{
			title: "操作",
			valueType: "option",
			width: 200,
			fixed: "right",
			render: (_, record) => (
				<Space>
					<HocAuth permission={["app:challenge:dailyStat:edit"]}>
						<LoadingButton key="edit" type="link" size="small" icon={<EditOutlined />} onClick={() => handleShowEdit(record)}>
							编辑
						</LoadingButton>
					</HocAuth>
					<HocAuth permission={["app:challenge:dailyStat:delete"]}>
						<LoadingButton
							key="del"
							danger
							type="link"
							size="small"
							icon={<DeleteOutlined />}
							onClick={() => handleDelete(record.statDate!)}
						>
							删除
						</LoadingButton>
					</HocAuth>
				</Space>
			)
		}
	];

	const toolBarRender = () => [
		<HocAuth key="add" permission={["app:challenge:dailyStat:add"]}>
			<LoadingButton type="primary" key="addBtn" icon={<PlusCircleOutlined />} onClick={() => handleShowCreate()}>
				新增
			</LoadingButton>
		</HocAuth>,
		<HocAuth key="export" permission={["app:challenge:dailyStat:export"]}>
			<LoadingButton type="primary" key="exportBtn" icon={<CloudDownloadOutlined />} onClick={done => handleExport(done)}>
				导出
			</LoadingButton>
		</HocAuth>
	];

	return (
		<>
			<ProTable<ChallengeDailyStatModel>
				rowKey="statDate"
				columns={columns}
				actionRef={actionRef}
				formRef={formRef}
				pagination={pagination}
				scroll={{ x: 1400 }}
				request={async params => {
					const { data } = await getDailyStatPageApi(params);
					return formatDataForProTable<ChallengeDailyStatModel>(data);
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

export default DailyStatPage;
