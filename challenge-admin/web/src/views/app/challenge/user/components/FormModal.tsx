import React, { forwardRef, useImperativeHandle, useState } from "react";
import { Modal, Form, Input, InputNumber, Select } from "antd";
import { addChallengeUserApi, updateChallengeUserApi, type ChallengeUserModel } from "@/api/app/challenge/user";
import { message } from "@/hooks/useMessage";

export interface FormModalRef {
	showAddFormModal: () => void;
	showEditFormModal: (id: number, record?: ChallengeUserModel) => void;
}

interface Props {
	onSuccess?: () => void;
}

const statusOptions = [
	{ label: "进行中", value: 1 },
	{ label: "成功", value: 2 },
	{ label: "失败", value: 3 }
];

const FormModal = forwardRef<FormModalRef, Props>(({ onSuccess }, ref) => {
	const [open, setOpen] = useState(false);
	const [confirmLoading, setConfirmLoading] = useState(false);
	const [isEdit, setIsEdit] = useState(false);
	const [currentId, setCurrentId] = useState<number>();
	const [form] = Form.useForm<ChallengeUserModel>();

	useImperativeHandle(ref, () => ({
		showAddFormModal: () => {
			setIsEdit(false);
			setCurrentId(undefined);
			form.resetFields();
			setOpen(true);
		},
		showEditFormModal: (id: number, record?: ChallengeUserModel) => {
			setIsEdit(true);
			setCurrentId(id);
			if (record) form.setFieldsValue(record);
			setOpen(true);
		}
	}));

	const handleOk = async () => {
		try {
			const values = await form.validateFields();
			setConfirmLoading(true);
			if (isEdit && currentId) {
				await updateChallengeUserApi(currentId, values);
				message.success("更新成功");
			} else {
				await addChallengeUserApi(values);
				message.success("新增成功");
			}
			setOpen(false);
			onSuccess?.();
		} catch (e) {
			// ignore
		} finally {
			setConfirmLoading(false);
		}
	};

	return (
		<Modal
			title={isEdit ? "编辑挑战用户" : "新增挑战用户"}
			open={open}
			onOk={handleOk}
			onCancel={() => setOpen(false)}
			confirmLoading={confirmLoading}
			destroyOnClose
		>
			<Form form={form} layout="vertical">
				<Form.Item label="用户ID" name="userId" rules={[{ required: true, message: "请输入用户ID" }]}>
					<InputNumber min={1} style={{ width: "100%" }} />
				</Form.Item>
				<Form.Item label="配置ID" name="configId">
					<InputNumber min={0} style={{ width: "100%" }} />
				</Form.Item>
				<Form.Item label="池子ID" name="poolId">
					<InputNumber min={0} style={{ width: "100%" }} />
				</Form.Item>
				<Form.Item label="挑战金额" name="challengeAmount" rules={[{ required: true, message: "请输入挑战金额" }]}>
					<Input />
				</Form.Item>
				<Form.Item label="开始日期" name="startDate">
					<InputNumber min={0} style={{ width: "100%" }} />
				</Form.Item>
				<Form.Item label="结束日期" name="endDate">
					<InputNumber min={0} style={{ width: "100%" }} />
				</Form.Item>
				<Form.Item label="状态" name="status" rules={[{ required: true, message: "请选择状态" }]}>
					<Select options={statusOptions} />
				</Form.Item>
				<Form.Item label="失败原因" name="failReason">
					<InputNumber min={0} style={{ width: "100%" }} />
				</Form.Item>
			</Form>
		</Modal>
	);
});

FormModal.displayName = "FormModal";

export default FormModal;
