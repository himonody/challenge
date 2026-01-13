import React, { forwardRef, useImperativeHandle, useState } from "react";
import { Modal, Form, Input, InputNumber, Select } from "antd";
import { addCheckinImageApi, updateCheckinImageApi, type ChallengeCheckinImageModel } from "@/api/app/challenge/checkinImage";
import { message } from "@/hooks/useMessage";

export interface FormModalRef {
	showAddFormModal: () => void;
	showEditFormModal: (id: number, record?: ChallengeCheckinImageModel) => void;
}

interface Props {
	onSuccess?: () => void;
}

const statusOptions = [
	{ label: "正常", value: 1 },
	{ label: "禁用", value: 0 }
];

// eslint-disable-next-line react/display-name
const FormModal = forwardRef<FormModalRef, Props>(({ onSuccess }, ref) => {
	const [open, setOpen] = useState(false);
	const [confirmLoading, setConfirmLoading] = useState(false);
	const [isEdit, setIsEdit] = useState(false);
	const [currentId, setCurrentId] = useState<number>();
	const [form] = Form.useForm<ChallengeCheckinImageModel>();

	useImperativeHandle(ref, () => ({
		showAddFormModal: () => {
			setIsEdit(false);
			setCurrentId(undefined);
			form.resetFields();
			setOpen(true);
		},
		showEditFormModal: (id: number, record?: ChallengeCheckinImageModel) => {
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
				await updateCheckinImageApi(currentId, values);
				message.success("更新成功");
			} else {
				await addCheckinImageApi(values);
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
			title={isEdit ? "编辑打卡图片" : "新增打卡图片"}
			open={open}
			onOk={handleOk}
			onCancel={() => setOpen(false)}
			confirmLoading={confirmLoading}
			destroyOnClose
		>
			<Form form={form} layout="vertical">
				<Form.Item label="打卡ID" name="checkinId" rules={[{ required: true, message: "请输入打卡ID" }]}>
					<InputNumber style={{ width: "100%" }} min={1} />
				</Form.Item>
				<Form.Item label="用户ID" name="userId" rules={[{ required: true, message: "请输入用户ID" }]}>
					<InputNumber style={{ width: "100%" }} min={1} />
				</Form.Item>
				<Form.Item label="图片URL" name="imageUrl" rules={[{ required: true, message: "请输入图片URL" }]}>
					<Input />
				</Form.Item>
				<Form.Item label="图片Hash" name="imageHash">
					<Input />
				</Form.Item>
				<Form.Item label="排序" name="sortNo">
					<InputNumber style={{ width: "100%" }} min={0} />
				</Form.Item>
				<Form.Item label="状态" name="status" rules={[{ required: true, message: "请选择状态" }]}>
					<Select options={statusOptions} />
				</Form.Item>
			</Form>
		</Modal>
	);
});

export default FormModal;
