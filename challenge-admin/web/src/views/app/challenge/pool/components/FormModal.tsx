import React, { forwardRef, useImperativeHandle, useState } from "react";
import { Modal, Form, Input, InputNumber, DatePicker, Select } from "antd";
import dayjs from "dayjs";
import { addPoolApi, updatePoolApi, type ChallengePoolModel } from "@/api/app/challenge/pool";
import { message } from "@/hooks/useMessage";

export interface FormModalRef {
	showAddFormModal: () => void;
	showEditFormModal: (id: number, record?: ChallengePoolModel) => void;
}

interface Props {
	onSuccess?: () => void;
}

type FormValues = Omit<ChallengePoolModel, "startDate" | "endDate"> & {
	startDate?: dayjs.Dayjs;
	endDate?: dayjs.Dayjs;
};

const settledOptions = [
	{ label: "未结算", value: 0 },
	{ label: "已结算", value: 1 }
];

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
		showEditFormModal: (id: number, record?: ChallengePoolModel) => {
			setIsEdit(true);
			setCurrentId(id);
			if (record) {
				form.setFieldsValue({
					...record,
					startDate: record.startDate ? dayjs(record.startDate) : undefined,
					endDate: record.endDate ? dayjs(record.endDate) : undefined
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
				startDate: values.startDate ? values.startDate.format("YYYY-MM-DD") : undefined,
				endDate: values.endDate ? values.endDate.format("YYYY-MM-DD") : undefined
			};
			if (isEdit && currentId) {
				await updatePoolApi(currentId, payload);
				message.success("更新成功");
			} else {
				await addPoolApi(payload);
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
			title={isEdit ? "编辑池子" : "新增池子"}
			open={open}
			onOk={handleOk}
			onCancel={() => setOpen(false)}
			confirmLoading={confirmLoading}
			destroyOnClose
		>
			<Form form={form} layout="vertical">
				<Form.Item label="配置ID" name="configId" rules={[{ required: true, message: "请输入配置ID" }]}>
					<InputNumber min={1} style={{ width: "100%" }} />
				</Form.Item>
				<Form.Item label="开始日期" name="startDate" rules={[{ required: true, message: "请选择开始日期" }]}>
					<DatePicker style={{ width: "100%" }} />
				</Form.Item>
				<Form.Item label="结束日期" name="endDate" rules={[{ required: true, message: "请选择结束日期" }]}>
					<DatePicker style={{ width: "100%" }} />
				</Form.Item>
				<Form.Item label="总金额" name="totalAmount" rules={[{ required: true, message: "请输入总金额" }]}>
					<Input />
				</Form.Item>
				<Form.Item label="是否结算" name="settled" rules={[{ required: true, message: "请选择结算状态" }]}>
					<Select options={settledOptions} />
				</Form.Item>
			</Form>
		</Modal>
	);
});

FormModal.displayName = "FormModal";

export default FormModal;
