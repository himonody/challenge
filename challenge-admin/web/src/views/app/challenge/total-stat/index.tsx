import { ExclamationCircleOutlined, PlusCircleOutlined, CloudDownloadOutlined, EditOutlined, DeleteOutlined } from "@ant-design/icons";
import { ProColumns, ProFormInstance, ProTable } from "@ant-design/pro-components";
import { Modal, Space } from "antd";
import React, { useRef } from "react";
import HocAuth from "@/components/HocAuth";
import LoadingButton from "@/components/LoadingButton";
import { pagination } from "@/config/proTable";
import { formatDataForProTable, saveExcelBlob } from "@/utils";
import { message, modal } from "@/hooks/useMessage";
import { ChallengeTotalStatModel, getTotalStatPageApi, exportTotalStatApi, delTotalStatApi } from "@/api/app/challenge/totalStat";
import type { ActionType } from "@ant-design/pro-components";
import FormModal from "./components/FormModal";
import type { FormModalRef } from "./components/FormModal";

const TotalStatPage: React.FC = () => {
	const actionRef = useRef<ActionType>();
	const formRef = useRef<ProFormInstance>();
	const formModalRef = useRef<FormModalRef>(null);

	const handleDelete = (id: number) => {
		Modal.confirm({
			title: "确认删除该总统计？",
			icon: <ExclamationCircleOutlined />,
			onOk: async () => {
				await delTotalStatApi(id);
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
					saveExcelBlob("challenge_total_stat", await exportTotalStatApi(formRef.current?.getFieldsValue()));
				} catch (err) {
					message.error("导出失败");
				} finally {
					done();
				}
			}
		});
	};

	const handleShowCreate = () => formModalRef.current?.showAddFormModal();
	const handleShowEdit = (record: ChallengeTotalStatModel) => {
		if (record.id === undefined) return;
		formModalRef.current?.showEditFormModal(record.id, record);
	};

	const columns: ProColumns<ChallengeTotalStatModel>[] = [
		{ title: "ID", dataIndex: "id", width: 80 },
		{ title: "总用户数", dataIndex: "totalUserCnt", width: 120 },
		{ title: "总参与次数", dataIndex: "totalJoinCnt", width: 120 },
		{ title: "总成功次数", dataIndex: "totalSuccessCnt", width: 120 },
		{ title: "总失败次数", dataIndex: "totalFailCnt", width: 120 },
		{ title: "总参与金额", dataIndex: "totalJoinAmount", width: 140, hideInSearch: true },
		{ title: "总成功金额", dataIndex: "totalSuccessAmount", width: 140, hideInSearch: true },
		{ title: "总失败金额", dataIndex: "totalFailAmount", width: 140, hideInSearch: true },
		{ title: "总平台奖励", dataIndex: "totalPlatformBonus", width: 140, hideInSearch: true },
		{ title: "总奖池金额", dataIndex: "totalPoolAmount", width: 140, hideInSearch: true },
		{ title: "更新时间", dataIndex: "updatedAt", valueType: "dateTime", width: 180, hideInSearch: true },
		{
			title: "操作",
			valueType: "option",
			width: 200,
			fixed: "right",
			render: (_, record) => (
				<Space>
					<HocAuth permission={["app:challenge:totalStat:edit"]}>
						<LoadingButton key="edit" type="link" size="small" icon={<EditOutlined />} onClick={() => handleShowEdit(record)}>
							编辑
						</LoadingButton>
					</HocAuth>
					<HocAuth permission={["app:challenge:totalStat:delete"]}>
						<LoadingButton key="del" danger type="link" size="small" icon={<DeleteOutlined />} onClick={() => handleDelete(record.id!)}>
							删除
						</LoadingButton>
					</HocAuth>
				</Space>
			)
		}
	];

	const toolBarRender = () => [
		<HocAuth key="add" permission={["app:challenge:totalStat:add"]}>
			<LoadingButton type="primary" key="addBtn" icon={<PlusCircleOutlined />} onClick={() => handleShowCreate()}>
				新增
			</LoadingButton>
		</HocAuth>,
		<HocAuth key="export" permission={["app:challenge:totalStat:export"]}>
			<LoadingButton type="primary" key="exportBtn" icon={<CloudDownloadOutlined />} onClick={done => handleExport(done)}>
				导出
			</LoadingButton>
		</HocAuth>
	];

	return (
		<>
			<ProTable<ChallengeTotalStatModel>
				rowKey="id"
				columns={columns}
				actionRef={actionRef}
				formRef={formRef}
				pagination={pagination}
				scroll={{ x: 1400 }}
				request={async params => {
					const { data } = await getTotalStatPageApi(params);
					return formatDataForProTable<ChallengeTotalStatModel>(data);
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

export default TotalStatPage;
