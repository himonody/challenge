import { ExclamationCircleOutlined, PlusCircleOutlined, CloudDownloadOutlined, EditOutlined, DeleteOutlined } from "@ant-design/icons";
import { ProColumns, ProFormInstance, ProTable } from "@ant-design/pro-components";
import { Modal, Space } from "antd";
import React, { useRef } from "react";
import HocAuth from "@/components/HocAuth";
import LoadingButton from "@/components/LoadingButton";
import { pagination } from "@/config/proTable";
import { formatDataForProTable, saveExcelBlob } from "@/utils";
import { message, modal } from "@/hooks/useMessage";
import { ChallengePoolModel, getPoolPageApi, exportPoolApi, delPoolApi } from "@/api/app/challenge/pool";
import type { ActionType } from "@ant-design/pro-components";
import FormModal from "./components/FormModal";
import type { FormModalRef } from "./components/FormModal";

const PoolPage: React.FC = () => {
	const actionRef = useRef<ActionType>();
	const formRef = useRef<ProFormInstance>();
	const formModalRef = useRef<FormModalRef>(null);

	const handleDelete = (id: number) => {
		Modal.confirm({
			title: "确认删除该池子？",
			icon: <ExclamationCircleOutlined />,
			onOk: async () => {
				await delPoolApi(id);
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
					saveExcelBlob("challenge_pool", await exportPoolApi(formRef.current?.getFieldsValue()));
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

	const handleShowEdit = (record: ChallengePoolModel) => {
		if (!record.id) return;
		formModalRef.current?.showEditFormModal(record.id, record);
	};

	const columns: ProColumns<ChallengePoolModel>[] = [
		{ title: "ID", dataIndex: "id", width: 80 },
		{ title: "配置ID", dataIndex: "configId", width: 100 },
		{ title: "开始日期", dataIndex: "startDate", valueType: "date", width: 120 },
		{ title: "结束日期", dataIndex: "endDate", valueType: "date", width: 120 },
		{ title: "总金额", dataIndex: "totalAmount", width: 140 },
		{
			title: "结算状态",
			dataIndex: "settled",
			valueType: "select",
			valueEnum: {
				0: { text: "未结算" },
				1: { text: "已结算" }
			},
			width: 100
		},
		{ title: "创建时间", dataIndex: "createdAt", valueType: "dateTime", width: 180, hideInSearch: true },
		{
			title: "操作",
			valueType: "option",
			width: 200,
			fixed: "right",
			render: (_, record) => (
				<Space>
					<HocAuth permission={["app:challenge:pool:edit"]}>
						<LoadingButton key="edit" type="link" size="small" icon={<EditOutlined />} onClick={() => handleShowEdit(record)}>
							编辑
						</LoadingButton>
					</HocAuth>
					<HocAuth permission={["app:challenge:pool:delete"]}>
						<LoadingButton key="del" danger type="link" size="small" icon={<DeleteOutlined />} onClick={() => handleDelete(record.id!)}>
							删除
						</LoadingButton>
					</HocAuth>
				</Space>
			)
		}
	];

	const toolBarRender = () => [
		<HocAuth key="add" permission={["app:challenge:pool:add"]}>
			<LoadingButton type="primary" key="addBtn" icon={<PlusCircleOutlined />} onClick={() => handleShowCreate()}>
				新增
			</LoadingButton>
		</HocAuth>,
		<HocAuth key="export" permission={["app:challenge:pool:export"]}>
			<LoadingButton type="primary" key="exportBtn" icon={<CloudDownloadOutlined />} onClick={done => handleExport(done)}>
				导出
			</LoadingButton>
		</HocAuth>
	];

	return (
		<>
			<ProTable<ChallengePoolModel>
				rowKey="id"
				columns={columns}
				actionRef={actionRef}
				formRef={formRef}
				pagination={pagination}
				scroll={{ x: 1200 }}
				request={async params => {
					const { data } = await getPoolPageApi(params);
					return formatDataForProTable<ChallengePoolModel>(data);
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

export default PoolPage;
