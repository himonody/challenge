import React, { forwardRef, useImperativeHandle, useState } from "react";
import { Modal, Form, DatePicker, Input, InputNumber, Select } from "antd";
import dayjs from "dayjs";
import type { ChallengeCheckinModel } from "@/api/app/challenge/checkin";
import { addCheckinApi, updateCheckinApi } from "@/api/app/challenge/checkin";
import { message } from "@/hooks/useMessage";

export interface FormModalRef {
	showAddFormModal: () => void;
	showEditFormModal: (id: number, record?: ChallengeCheckinModel) => void;
}

interface Props {
	onSuccess?: () => void;
}

const statusOptions = [
	{ label: "正常", value: 1 },
	{ label: "禁用", value: 0 }
];

type FormValues = Omit<ChallengeCheckinModel, "checkinDate" | "checkinTime"> & {
	checkinDate?: dayjs.Dayjs;
	checkinTime?: dayjs.Dayjs;
};

const FormModal = forwardRef<FormModalRef, Props>(({ onSuccess }, ref) => {
	const [open, setOpen] = useState(false);
	const [confirmLoading, setConfirmLoading] = useState(false);
	const [isEdit, setIsEdit] = useState(false);
	const [currentId, setCurrentId] = useState<number>();
	const [form] = Form.useForm<FormValues>();

	useImperativeHandle(ref, () => ({
		showAddFormModal: () => {
			setIsEdit(false);
			setCurrentId(undefined);
			form.resetFields();
			setOpen(true);
		},
		showEditFormModal: (id: number, record?: ChallengeCheckinModel) => {
			setIsEdit(true);
			setCurrentId(id);
			if (record) {
				form.setFieldsValue({
					...record,
					checkinDate: record.checkinDate ? dayjs(record.checkinDate) : undefined,
					checkinTime: record.checkinTime ? dayjs(record.checkinTime) : undefined
				});
			}
			setOpen(true);
		}
	}));

	const handleOk = async () => {
		try {
			const values = await form.validateFields();
			setConfirmLoading(true);
			const payload = {
				...values,
				checkinDate: values.checkinDate ? values.checkinDate.format("YYYY-MM-DD") : undefined,
				checkinTime: values.checkinTime ? values.checkinTime.format("YYYY-MM-DD HH:mm:ss") : undefined
			};
			if (isEdit && currentId) {
				await updateCheckinApi(currentId, payload);
				message.success("更新成功");
			} else {
				await addCheckinApi(payload);
				message.success("新增成功");
			}
			setOpen(false);
			onSuccess?.();
		} catch (err) {
			// 校验错误不提示
		} finally {
			setConfirmLoading(false);
		}
	};

	return (
		<Modal
			title={isEdit ? "编辑打卡" : "新增打卡"}
			open={open}
			onOk={handleOk}
			onCancel={() => setOpen(false)}
			confirmLoading={confirmLoading}
			destroyOnClose
		>
			<Form form={form} layout="vertical">
				<Form.Item label="挑战ID" name="challengeId" rules={[{ required: true, message: "请输入挑战ID" }]}>
					<InputNumber style={{ width: "100%" }} min={1} />
				</Form.Item>
				<Form.Item label="用户ID" name="userId" rules={[{ required: true, message: "请输入用户ID" }]}>
					<InputNumber style={{ width: "100%" }} min={1} />
				</Form.Item>
				<Form.Item label="打卡日期" name="checkinDate" rules={[{ required: true, message: "请选择打卡日期" }]}>
					<DatePicker style={{ width: "100%" }} />
				</Form.Item>
				<Form.Item label="打卡时间" name="checkinTime">
					<DatePicker showTime style={{ width: "100%" }} />
				</Form.Item>
				<Form.Item label="心情代码" name="moodCode">
					<InputNumber style={{ width: "100%" }} min={0} />
				</Form.Item>
				<Form.Item label="心情描述" name="moodText">
					<Input />
				</Form.Item>
				<Form.Item label="内容类型" name="contentType">
					<InputNumber style={{ width: "100%" }} min={0} />
				</Form.Item>
				<Form.Item label="状态" name="status" rules={[{ required: true, message: "请选择状态" }]}>
					<Select options={statusOptions} />
				</Form.Item>
			</Form>
		</Modal>
	);
});

FormModal.displayName = "FormModal";

export default FormModal;
