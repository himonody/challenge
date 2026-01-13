import React, { forwardRef, useImperativeHandle, useState } from "react";
import { Modal, Form, Input, InputNumber, Select } from "antd";
import { addPoolFlowApi, updatePoolFlowApi, type ChallengePoolFlowModel } from "@/api/app/challenge/poolFlow";
import { message } from "@/hooks/useMessage";

export interface FormModalRef {
	showAddFormModal: () => void;
	showEditFormModal: (id: number, record?: ChallengePoolFlowModel) => void;
}

interface Props {
	onSuccess?: () => void;
}

const typeOptions = [
	{ label: "入账", value: 1 },
	{ label: "出账", value: 2 }
];

const FormModal = forwardRef<FormModalRef, Props>(({ onSuccess }, ref) => {
	const [open, setOpen] = useState(false);
	const [confirmLoading, setConfirmLoading] = useState(false);
	const [isEdit, setIsEdit] = useState(false);
	const [currentId, setCurrentId] = useState<number>();
	const [form] = Form.useForm<ChallengePoolFlowModel>();

	useImperativeHandle(ref, () => ({
		showAddFormModal: () => {
			setIsEdit(false);
			setCurrentId(undefined);
			form.resetFields();
			setOpen(true);
		},
		showEditFormModal: (id: number, record?: ChallengePoolFlowModel) => {
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
				await updatePoolFlowApi(currentId, values);
				message.success("更新成功");
			} else {
				await addPoolFlowApi(values);
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
			title={isEdit ? "编辑池子流水" : "新增池子流水"}
			open={open}
			onOk={handleOk}
			onCancel={() => setOpen(false)}
			confirmLoading={confirmLoading}
			destroyOnClose
		>
			<Form form={form} layout="vertical">
				<Form.Item label="池子ID" name="poolId" rules={[{ required: true, message: "请输入池子ID" }]}>
					<InputNumber min={1} style={{ width: "100%" }} />
				</Form.Item>
				<Form.Item label="用户ID" name="userId">
					<InputNumber min={0} style={{ width: "100%" }} />
				</Form.Item>
				<Form.Item label="金额" name="amount" rules={[{ required: true, message: "请输入金额" }]}>
					<Input />
				</Form.Item>
				<Form.Item label="类型" name="type" rules={[{ required: true, message: "请选择类型" }]}>
					<Select options={typeOptions} />
				</Form.Item>
			</Form>
		</Modal>
	);
});

FormModal.displayName = "FormModal";

export default FormModal;
