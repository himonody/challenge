import { getDictOptions, getDictsApi } from "@/api/admin/sys/sys-dictdata";
import {
	exportUserApi,
	getUserPageApi,
	UserModel,
	rechargeUserApi,
	deductUserApi,
	resetUserPasswordApi,
	resetUserPayPasswordApi,
	updateUserPayStatusApi
} from "@/api/app/user/user";
import HocAuth from "@/components/HocAuth";
import LoadingButton from "@/components/LoadingButton";
import { pagination } from "@/config/proTable";
import { SummaryColor } from "@/enums/base";
import { ResultEnum } from "@/enums/httpEnum";
import { message, modal } from "@/hooks/useMessage";
import { formatDataForProTable, saveExcelBlob } from "@/utils";
import { CloudDownloadOutlined, ExclamationCircleOutlined, PlusCircleOutlined } from "@ant-design/icons";
import type { ActionType, ProColumns, ProFormInstance } from "@ant-design/pro-components";
import { ProTable } from "@ant-design/pro-components";
import { Statistic, Modal, InputNumber, Select } from "antd";
import React, { useEffect, useRef, useState } from "react";
import FormModal, { FormModalRef } from "./components/FormModal";

const User: React.FC = () => {
	const actionRef = React.useRef<ActionType>();
	const tableFormRef = React.useRef<ProFormInstance>();
	const formModalRef = useRef<FormModalRef>(null);
	const [statusOptions, setStatusOptions] = useState<Map<string, string>>(new Map());
	const [levelTypeOptions, setLevelTypeOptions] = useState<Map<string, string>>(new Map());
	const [extend, setExtend] = useState<UserModel>({});

	// 操作函数
	const handleRecharge = (id: number) => {
		let amount = 0;
		Modal.confirm({
			title: "人工充值",
			content: (
				<InputNumber
					style={{ width: "100%" }}
					min={0}
					precision={2}
					onChange={value => (amount = Number(value))}
					placeholder="请输入充值金额"
				/>
			),
			onOk: async () => {
				if (amount <= 0) {
					message.error("金额必须大于0");
					throw new Error("invalid amount");
				}
				await rechargeUserApi(id, { amount });
				message.success("充值成功");
				actionRef.current?.reload();
			}
		});
	};

	const handleDeduct = (id: number) => {
		let amount = 0;
		Modal.confirm({
			title: "人工扣款",
			content: (
				<InputNumber
					style={{ width: "100%" }}
					min={0}
					precision={2}
					onChange={value => (amount = Number(value))}
					placeholder="请输入扣款金额"
				/>
			),
			onOk: async () => {
				if (amount <= 0) {
					message.error("金额必须大于0");
					throw new Error("invalid amount");
				}
				await deductUserApi(id, { amount });
				message.success("扣款成功");
				actionRef.current?.reload();
			}
		});
	};

	const handleResetPwd = async (id: number) => {
		await resetUserPasswordApi(id);
		message.success("已重置登录密码为:aa1234");
	};

	const handleResetPayPwd = async (id: number) => {
		await resetUserPayPasswordApi(id);
		message.success("已重置支付密码为:aa1234");
	};

	const handleUpdatePayStatus = (id: number) => {
		let status = "1";
		Modal.confirm({
			title: "支付状态变更",
			content: (
				<Select
					defaultValue="1"
					style={{ width: "100%" }}
					onChange={value => (status = value)}
					options={[
						{ label: "启用", value: "1" },
						{ label: "禁用", value: "2" }
					]}
				/>
			),
			onOk: async () => {
				await updateUserPayStatusApi(id, { pay_status: status });
				message.success("支付状态已更新");
				actionRef.current?.reload();
			}
		});
	};

	// 定义列
	const columns: ProColumns<UserModel>[] = [
		{
			title: "序号",
			dataIndex: "index",
			valueType: "index",
			width: 50,
			align: "center",
			className: "gray-cell",
			render: (_, __, index, action) => {
				// 根据分页计算实际序号
				const currentPage = action?.pageInfo?.current || 1;
				const pageSize = action?.pageInfo?.pageSize || 10;
				return (currentPage - 1) * pageSize + index + 1;
			}
		},
		{
			title: "用户编号",
			dataIndex: "id",
			width: 80,
			align: "left"
		},
		{
			title: "用户名",
			dataIndex: "userName",
			width: 120,
			align: "left"
		},
		{
			title: "用户昵称",
			dataIndex: "nickName",
			width: 120,
			align: "left",
			hideInSearch: true
		},
		{
			title: "等级编号",
			dataIndex: "levelId",
			width: 80,
			align: "left",
			hideInSearch: true,
			hideInTable: true
		},
		{
			title: "真实姓名",
			dataIndex: "trueName",
			width: 80,
			align: "left",
			hideInTable: true
		},
		{
			title: "账户余额",
			dataIndex: "money",
			hideInSearch: true,
			width: 80,
			align: "left"
		},
		{
			title: "冻结金额",
			dataIndex: "freezeMoney",
			hideInSearch: true,
			width: 90,
			align: "left"
		},
		{
			title: "电子邮箱",
			dataIndex: "email",
			width: 180,
			align: "left",
			hideInTable: true
		},
		{
			title: "国家区号",
			dataIndex: "mobileTitle",
			hideInSearch: true,
			width: 80,
			align: "left",
			hideInTable: true
		},
		{
			title: "手机号码",
			dataIndex: "mobile",
			width: 120,
			align: "left",
			hideInTable: true
		},
		{
			title: "等级类型",
			dataIndex: "levelType",
			valueType: "select",
			width: 80,
			align: "left",
			hideInSearch: true,
			hideInTable: true,
			valueEnum: Object.fromEntries(levelTypeOptions)
		},
		{
			title: "等级",
			dataIndex: "level",
			width: 80,
			align: "left",
			hideInSearch: true,
			hideInTable: true,
			render: (text, record) => record.userLevel?.level
		},
		{
			title: "当前用户邀请码",
			dataIndex: "refCode",
			width: 120,
			align: "left"
		},
		{
			title: "上级用户邀请码",
			dataIndex: "parentRefCode",
			width: 120,
			align: "left"
		},
		{
			title: "上级用户编号",
			dataIndex: "parentId",
			width: 120,
			align: "left",
			hideInSearch: true
		},
		{
			title: "注册时间",
			dataIndex: "registerAt",
			valueType: "dateTime",
			hideInSearch: true,
			width: 180,
			align: "left"
		},
		{
			title: "注册IP",
			dataIndex: "registerIp",
			hideInSearch: true,
			width: 140,
			align: "left"
		},
		{
			title: "最后登录时间",
			dataIndex: "lastLoginAt",
			valueType: "dateTime",
			hideInSearch: true,
			width: 180,
			align: "left"
		},
		{
			title: "最后登录IP",
			dataIndex: "lastLoginIp",
			hideInSearch: true,
			width: 140,
			align: "left"
		},
		{
			title: "状态",
			dataIndex: "status",
			valueType: "select",
			valueEnum: statusOptions,
			width: 80,
			align: "left"
		},
		{
			title: "创建时间",
			dataIndex: "createdAt",
			hideInSearch: true,
			valueType: "dateTime",
			width: 180,
			align: "left"
		},
		{
			title: "创建时间",
			dataIndex: "createdAt",
			valueType: "dateTimeRange",
			hideInTable: true,
			search: { transform: value => ({ beginCreatedAt: value[0], endCreatedAt: value[1] }) }
		},
		{
			title: "操作",
			valueType: "option",
			align: "center",
			fixed: "right",
			width: 400,
			render: (_, data) => (
				<HocAuth permission={["app:user:edit"]}>
					<div style={{ display: "grid", gridTemplateColumns: "repeat(3, auto)", gap: 8 }}>
						<LoadingButton key="edit" type="link" size="small" onClick={done => handleShowEditFormModal(data.id!, done)}>
							编辑
						</LoadingButton>
						<LoadingButton
							key="recharge"
							type="link"
							size="small"
							onClick={done => {
								handleRecharge(data.id!);
								done();
							}}
							danger={false}
						>
							充值
						</LoadingButton>
						<LoadingButton
							key="deduct"
							type="link"
							size="small"
							onClick={done => {
								handleDeduct(data.id!);
								done();
							}}
							danger
						>
							扣款
						</LoadingButton>
						<LoadingButton
							key="resetPwd"
							type="link"
							size="small"
							onClick={async done => {
								await handleResetPwd(data.id!);
								done();
							}}
						>
							重置登录密码
						</LoadingButton>
						<LoadingButton
							key="resetPayPwd"
							type="link"
							size="small"
							onClick={async done => {
								await handleResetPayPwd(data.id!);
								done();
							}}
						>
							重置支付密码
						</LoadingButton>
						<LoadingButton
							key="payStatus"
							type="link"
							size="small"
							onClick={done => {
								handleUpdatePayStatus(data.id!);
								done();
							}}
						>
							支付状态
						</LoadingButton>
					</div>
				</HocAuth>
			)
		}
	];

	useEffect(() => {
		const initData = async () => {
			const { data: statusData, msg: statusMsg, code: statusCode } = await getDictsApi("admin_sys_status");
			if (statusCode !== ResultEnum.SUCCESS) {
				message.error(statusMsg);
				return;
			}
			setStatusOptions(getDictOptions(statusData));
			const { data: levelTypeData, msg: levelTypeMsg, code: levelTypeCode } = await getDictsApi("app_user_level_type");
			if (levelTypeCode !== ResultEnum.SUCCESS) {
				message.error(levelTypeMsg);
				return;
			}
			setLevelTypeOptions(getDictOptions(levelTypeData));
		};
		initData();
	}, []);

	const handleShowAddFormModal = (done: () => void) => {
		formModalRef.current?.showAddFormModal();
		setTimeout(() => done(), 1000);
	};

	const handleShowEditFormModal = (id: number, done: () => void) => {
		formModalRef.current?.showEditFormModal(id);
		setTimeout(() => done(), 1000);
	};

	const handleFormModalConfirm = () => {
		actionRef.current?.reload(false);
	};

	const handleExport = (done: () => void) => {
		modal.confirm({
			title: "提示",
			icon: <ExclamationCircleOutlined />,
			content: "是否确认导出所选数据？",
			okText: "确认",
			cancelText: "取消",
			maskClosable: true,
			onCancel: () => {
				done();
			},
			onOk: async () => {
				try {
					saveExcelBlob("用户管理", await exportUserApi(tableFormRef.current?.getFieldsValue()));
				} catch (err) {
					message.error("下载失败，请检查网络");
				} finally {
					done();
				}
			}
		});
	};

	const toolBarRender = () => [
		<HocAuth key="addAuth" permission={["app:user:add"]}>
			<LoadingButton type="primary" key="addTable" icon={<PlusCircleOutlined />} onClick={done => handleShowAddFormModal(done)}>
				新增
			</LoadingButton>
		</HocAuth>,
		<HocAuth key="exportAuth" permission={["app:user:export"]}>
			<LoadingButton type="primary" key="importTable" icon={<CloudDownloadOutlined />} onClick={done => handleExport(done)}>
				Excel导出
			</LoadingButton>
		</HocAuth>
	];

	return (
		<>
			<ProTable<UserModel>
				className="ant-pro-table-scroll"
				columns={columns}
				actionRef={actionRef}
				formRef={tableFormRef}
				bordered
				cardBordered
				defaultSize="small"
				scroll={{ x: "2000", y: "100%" }}
				request={async params => {
					const { data } = await getUserPageApi(params);
					setExtend(data.extend);
					return formatDataForProTable<UserModel>(data);
				}}
				columnsState={{
					persistenceKey: "use-pro-table-key",
					persistenceType: "localStorage"
				}}
				options={{
					reload: true,
					density: true,
					fullScreen: true
				}}
				rowKey="id"
				search={{ labelWidth: "auto", showHiddenNum: true }}
				pagination={pagination}
				dateFormatter="string"
				headerTitle="用户管理"
				toolBarRender={toolBarRender}
				footer={() => extend && <Statistic title="余额 总计" value={extend.money} valueStyle={{ color: SummaryColor.base }} />}
			/>
			<FormModal ref={formModalRef} onConfirm={handleFormModalConfirm} />
		</>
	);
};

export default User;
