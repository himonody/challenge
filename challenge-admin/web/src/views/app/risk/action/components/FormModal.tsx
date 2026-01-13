import React, { forwardRef, useImperativeHandle, useState } from "react";
import { Modal, Form, Input, InputNumber } from "antd";
import { addRiskActionApi, updateRiskActionApi, type RiskActionModel } from "@/api/app/risk/action";
import { message } from "@/hooks/useMessage";

export interface FormModalRef {
	showAddFormModal: () => void;
	showEditFormModal: (code: string, record?: RiskActionModel) => void;
}

interface Props {
	onSuccess?: () => void;
}

const FormModal = forwardRef<FormModalRef, Props>(({ onSuccess }, ref) => {
	const [open, setOpen] = useState(false);
	const [confirmLoading, setConfirmLoading] = useState(false);
	const [isEdit, setIsEdit] = useState(false);
	const [currentCode, setCurrentCode] = useState<string>();
	const [form] = Form.useForm<RiskActionModel>();

	useImperativeHandle(ref, () => ({
		showAddFormModal: () => {
			setIsEdit(false);
			setCurrentCode(undefined);
			form.resetFields();
			setOpen(true);
		},
		showEditFormModal: (code: string, record?: RiskActionModel) => {
			setIsEdit(true);
			setCurrentCode(code);
			if (record) form.setFieldsValue(record);
			setOpen(true);
		}
	}));

	const handleOk = async () => {
		try {
			const values = await form.validateFields();
			setConfirmLoading(true);
			if (isEdit && currentCode) {
				await updateRiskActionApi(currentCode, values);
				message.success("更新成功");
			} else {
				await addRiskActionApi(values);
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
			title={isEdit ? "编辑动作" : "新增动作"}
			open={open}
			onOk={handleOk}
			onCancel={() => setOpen(false)}
			confirmLoading={confirmLoading}
			destroyOnClose
		>
			<Form form={form} layout="vertical">
				<Form.Item label="动作编码" name="code" rules={[{ required: true, message: "请输入动作编码" }]}>
					<Input disabled={isEdit} />
				</Form.Item>
				<Form.Item label="动作类型" name="type" rules={[{ required: true, message: "请输入动作类型" }]}>
					<Input />
				</Form.Item>
				<Form.Item label="默认动作值" name="defaultValue" rules={[{ required: true, message: "请输入默认动作值" }]}>
					<InputNumber min={0} style={{ width: "100%" }} />
				</Form.Item>
				<Form.Item label="说明" name="remark">
					<Input.TextArea rows={3} />
				</Form.Item>
			</Form>
		</Modal>
	);
});

FormModal.displayName = "FormModal";

export default FormModal;
