import React, { forwardRef, useImperativeHandle, useState } from "react";
import { Modal, Form, Input, InputNumber, Select, DatePicker } from "antd";
import dayjs from "dayjs";
import { message } from "@/hooks/useMessage";
import { exportRiskRateLimitApi } from "@/api/app/risk/rateLimit";

export interface FormModalRef {
	showAddFormModal: () => void;
}

interface Props {
	onExport?: () => void;
}

const FormModal = forwardRef<FormModalRef, Props>(({ onExport }, ref) => {
	const [open, setOpen] = useState(false);
	const [confirmLoading, setConfirmLoading] = useState(false);
	const [form] = Form.useForm();

	useImperativeHandle(ref, () => ({
		showAddFormModal: () => {
			form.resetFields();
			setOpen(true);
		}
	}));

	const handleOk = async () => {
		try {
			const values = await form.validateFields();
			const query = {
				...values,
				beginCreatedAt: values.windowStart ? values.windowStart.format("YYYY-MM-DD HH:mm:ss") : undefined,
				endCreatedAt: values.windowEnd ? values.windowEnd.format("YYYY-MM-DD HH:mm:ss") : undefined
			};
			setConfirmLoading(true);
			await exportRiskRateLimitApi(query);
			message.success("导出已开始");
			onExport?.();
			setOpen(false);
		} catch (e) {
			// ignore
		} finally {
			setConfirmLoading(false);
		}
	};

	return (
		<Modal title="导出限流记录" open={open} onOk={handleOk} onCancel={() => setOpen(false)} confirmLoading={confirmLoading} destroyOnClose>
			<Form form={form} layout="vertical">
				<Form.Item label="场景" name="scene">
					<Input />
				</Form.Item>
				<Form.Item label="标识类型" name="identityType">
					<Input />
				</Form.Item>
				<Form.Item label="标识值" name="identityValue">
					<Input />
				</Form.Item>
				<Form.Item label="窗口开始" name="windowStart">
					<DatePicker showTime style={{ width: "100%" }} />
				</Form.Item>
				<Form.Item label="窗口结束" name="windowEnd">
					<DatePicker showTime style={{ width: "100%" }} />
				</Form.Item>
				<Form.Item label="是否拦截" name="blocked">
					<Select
						allowClear
						options={[
							{ label: "是", value: "1" },
							{ label: "否", value: "0" }
						]}
					/>
				</Form.Item>
			</Form>
		</Modal>
	);
});

FormModal.displayName = "FormModal";

export default FormModal;
