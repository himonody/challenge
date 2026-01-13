import { ExclamationCircleOutlined, PlusCircleOutlined, CloudDownloadOutlined, EditOutlined, DeleteOutlined } from "@ant-design/icons";
import { ProColumns, ProFormInstance, ProTable } from "@ant-design/pro-components";
import { Modal, Space } from "antd";
import React, { useRef } from "react";
import HocAuth from "@/components/HocAuth";
import LoadingButton from "@/components/LoadingButton";
import { pagination } from "@/config/proTable";
import { formatDataForProTable, saveExcelBlob } from "@/utils";
import { message, modal } from "@/hooks/useMessage";
import {
	ChallengeCheckinVideoAdModel,
	getCheckinVideoAdPageApi,
	exportCheckinVideoAdApi,
	delCheckinVideoAdApi
} from "@/api/app/challenge/checkinVideoAd";
import type { ActionType } from "@ant-design/pro-components";
import FormModal from "./components/FormModal";
import type { FormModalRef } from "./components/FormModal";

const CheckinVideoAd: React.FC = () => {
	const actionRef = useRef<ActionType>();
	const formRef = useRef<ProFormInstance>();
	const formModalRef = useRef<FormModalRef>(null);

	const handleDelete = (id: number) => {
		Modal.confirm({
			title: "确认删除该视频广告记录？",
			icon: <ExclamationCircleOutlined />,
			onOk: async () => {
				await delCheckinVideoAdApi(id);
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
					saveExcelBlob("challenge_checkin_video_ad", await exportCheckinVideoAdApi(formRef.current?.getFieldsValue()));
				} catch (err) {
					message.error("导出失败");
				} finally {
					done();
				}
			}
		});
	};

	const handleShowCreate = () => {
		formModalRef.current?.showAddFormModal();
	};

	const handleShowEdit = (record: ChallengeCheckinVideoAdModel) => {
		if (!record.id) return;
		formModalRef.current?.showEditFormModal(record.id, record);
	};

	const columns: ProColumns<ChallengeCheckinVideoAdModel>[] = [
		{ title: "ID", dataIndex: "id", width: 80 },
		{ title: "打卡ID", dataIndex: "checkinId", width: 120 },
		{ title: "用户ID", dataIndex: "userId", width: 120 },
		{ title: "广告平台", dataIndex: "adPlatform", width: 120 },
		{ title: "广告位ID", dataIndex: "adUnitId", width: 160, hideInSearch: true },
		{ title: "广告订单号", dataIndex: "adOrderNo", width: 160, hideInSearch: true },
		{ title: "视频时长(秒)", dataIndex: "videoDuration", width: 110, hideInSearch: true },
		{ title: "观看时长(秒)", dataIndex: "watchDuration", width: 110, hideInSearch: true },
		{ title: "收益金额", dataIndex: "rewardAmount", width: 120, hideInSearch: true },
		{ title: "校验状态", dataIndex: "verifyStatus", width: 100 },
		{ title: "观看完成时间", dataIndex: "createdAt", valueType: "dateTime", width: 180, hideInSearch: true },
		{ title: "校验完成时间", dataIndex: "verifiedAt", valueType: "dateTime", width: 180, hideInSearch: true },
		{
			title: "操作",
			valueType: "option",
			width: 200,
			fixed: "right",
			render: (_, record) => (
				<Space>
					<HocAuth permission={["app:challenge:checkinVideoAd:edit"]}>
						<LoadingButton key="edit" type="link" size="small" icon={<EditOutlined />} onClick={() => handleShowEdit(record)}>
							编辑
						</LoadingButton>
					</HocAuth>
					<HocAuth permission={["app:challenge:checkinVideoAd:delete"]}>
						<LoadingButton key="del" danger type="link" size="small" icon={<DeleteOutlined />} onClick={() => handleDelete(record.id!)}>
							删除
						</LoadingButton>
					</HocAuth>
				</Space>
			)
		}
	];

	const toolBarRender = () => [
		<HocAuth key="add" permission={["app:challenge:checkinVideoAd:add"]}>
			<LoadingButton type="primary" key="addBtn" icon={<PlusCircleOutlined />} onClick={() => handleShowCreate()}>
				新增
			</LoadingButton>
		</HocAuth>,
		<HocAuth key="export" permission={["app:challenge:checkinVideoAd:export"]}>
			<LoadingButton type="primary" key="exportBtn" icon={<CloudDownloadOutlined />} onClick={done => handleExport(done)}>
				导出
			</LoadingButton>
		</HocAuth>
	];

	return (
		<>
			<ProTable<ChallengeCheckinVideoAdModel>
				rowKey="id"
				columns={columns}
				actionRef={actionRef}
				formRef={formRef}
				pagination={pagination}
				scroll={{ x: 1400 }}
				request={async params => {
					const { data } = await getCheckinVideoAdPageApi(params);
					return formatDataForProTable<ChallengeCheckinVideoAdModel>(data);
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

export default CheckinVideoAd;
