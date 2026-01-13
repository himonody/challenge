import { ExclamationCircleOutlined, PlusCircleOutlined, CloudDownloadOutlined, EditOutlined, DeleteOutlined } from "@ant-design/icons";
import { ProColumns, ProFormInstance, ProTable } from "@ant-design/pro-components";
import { Modal, Space } from "antd";
import React, { useRef } from "react";
import HocAuth from "@/components/HocAuth";
import LoadingButton from "@/components/LoadingButton";
import { pagination } from "@/config/proTable";
import { formatDataForProTable, saveExcelBlob } from "@/utils";
import { message, modal } from "@/hooks/useMessage";
import { ChallengePoolFlowModel, getPoolFlowPageApi, exportPoolFlowApi, delPoolFlowApi } from "@/api/app/challenge/poolFlow";
import type { ActionType } from "@ant-design/pro-components";
import FormModal from "./components/FormModal";
import type { FormModalRef } from "./components/FormModal";

const PoolFlowPage: React.FC = () => {
	const actionRef = useRef<ActionType>();
	const formRef = useRef<ProFormInstance>();
	const formModalRef = useRef<FormModalRef>(null);

	const handleDelete = (id: number) => {
		Modal.confirm({
			title: "确认删除该流水？",
			icon: <ExclamationCircleOutlined />,
			onOk: async () => {
				await delPoolFlowApi(id);
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
					saveExcelBlob("challenge_pool_flow", await exportPoolFlowApi(formRef.current?.getFieldsValue()));
				} catch (err) {
					message.error("导出失败");
				} finally {
					done();
				}
			}
		});
	};

	const handleShowCreate = () => formModalRef.current?.showAddFormModal();
	const handleShowEdit = (record: ChallengePoolFlowModel) => {
		if (!record.id) return;
		formModalRef.current?.showEditFormModal(record.id, record);
	};

	const columns: ProColumns<ChallengePoolFlowModel>[] = [
		{ title: "ID", dataIndex: "id", width: 80 },
		{ title: "池子ID", dataIndex: "poolId", width: 100 },
		{ title: "用户ID", dataIndex: "userId", width: 100 },
		{ title: "金额", dataIndex: "amount", width: 120 },
		{
			title: "类型",
			dataIndex: "type",
			valueType: "select",
			valueEnum: {
				1: { text: "入账" },
				2: { text: "出账" }
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
					<HocAuth permission={["app:challenge:poolFlow:edit"]}>
						<LoadingButton key="edit" type="link" size="small" icon={<EditOutlined />} onClick={() => handleShowEdit(record)}>
							编辑
						</LoadingButton>
					</HocAuth>
					<HocAuth permission={["app:challenge:poolFlow:delete"]}>
						<LoadingButton key="del" danger type="link" size="small" icon={<DeleteOutlined />} onClick={() => handleDelete(record.id!)}>
							删除
						</LoadingButton>
					</HocAuth>
				</Space>
			)
		}
	];

	const toolBarRender = () => [
		<HocAuth key="add" permission={["app:challenge:poolFlow:add"]}>
			<LoadingButton type="primary" key="addBtn" icon={<PlusCircleOutlined />} onClick={() => handleShowCreate()}>
				新增
			</LoadingButton>
		</HocAuth>,
		<HocAuth key="export" permission={["app:challenge:poolFlow:export"]}>
			<LoadingButton type="primary" key="exportBtn" icon={<CloudDownloadOutlined />} onClick={done => handleExport(done)}>
				导出
			</LoadingButton>
		</HocAuth>
	];

	return (
		<>
			<ProTable<ChallengePoolFlowModel>
				rowKey="id"
				columns={columns}
				actionRef={actionRef}
				formRef={formRef}
				pagination={pagination}
				scroll={{ x: 1200 }}
				request={async params => {
					const { data } = await getPoolFlowPageApi(params);
					return formatDataForProTable<ChallengePoolFlowModel>(data);
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

export default PoolFlowPage;
