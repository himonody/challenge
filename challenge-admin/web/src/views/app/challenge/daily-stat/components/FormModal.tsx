import React, { forwardRef, useImperativeHandle, useState } from "react";
import { Modal, Form, Input, InputNumber, DatePicker } from "antd";
import dayjs from "dayjs";
import { addDailyStatApi, updateDailyStatApi, type ChallengeDailyStatModel } from "@/api/app/challenge/dailyStat";
import { message } from "@/hooks/useMessage";

export interface FormModalRef {
	showAddFormModal: () => void;
	showEditFormModal: (statDate: string, record?: ChallengeDailyStatModel) => void;
}

interface Props {
	onSuccess?: () => void;
}

type FormValues = Omit<ChallengeDailyStatModel, "statDate"> & {
	statDate?: dayjs.Dayjs;
};

const FormModal = forwardRef<FormModalRef, Props>(({ onSuccess }, ref) => {
	const [open, setOpen] = useState(false);
	const [confirmLoading, setConfirmLoading] = useState(false);
	const [isEdit, setIsEdit] = useState(false);
	const [currentStatDate, setCurrentStatDate] = useState<string>();
	const [form] = Form.useForm<FormValues>();

	useImperativeHandle(ref, () => ({
		showAddFormModal: () => {
			setIsEdit(false);
			setCurrentStatDate(undefined);
			form.resetFields();
			setOpen(true);
		},
		showEditFormModal: (statDate: string, record?: ChallengeDailyStatModel) => {
			setIsEdit(true);
			setCurrentStatDate(statDate);
			if (record) {
				form.setFieldsValue({
					...record,
					statDate: record.statDate ? dayjs(record.statDate) : undefined
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
				statDate: values.statDate ? values.statDate.format("YYYY-MM-DD") : currentStatDate
			};
			if (isEdit && currentStatDate) {
				await updateDailyStatApi(currentStatDate, payload);
				message.success("更新成功");
			} else {
				await addDailyStatApi(payload);
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
			title={isEdit ? "编辑每日统计" : "新增每日统计"}
			open={open}
			onOk={handleOk}
			onCancel={() => setOpen(false)}
			confirmLoading={confirmLoading}
			destroyOnClose
		>
			<Form form={form} layout="vertical">
				<Form.Item label="统计日期" name="statDate" rules={[{ required: true, message: "请选择统计日期" }]}>
					<DatePicker style={{ width: "100%" }} disabled={isEdit} />
				</Form.Item>
				<Form.Item label="参与人数" name="joinUserCnt">
					<InputNumber min={0} style={{ width: "100%" }} />
				</Form.Item>
				<Form.Item label="成功人数" name="successUserCnt">
					<InputNumber min={0} style={{ width: "100%" }} />
				</Form.Item>
				<Form.Item label="失败人数" name="failUserCnt">
					<InputNumber min={0} style={{ width: "100%" }} />
				</Form.Item>
				<Form.Item label="参与金额" name="joinAmount">
					<Input />
				</Form.Item>
				<Form.Item label="成功金额" name="successAmount">
					<Input />
				</Form.Item>
				<Form.Item label="失败金额" name="failAmount">
					<Input />
				</Form.Item>
				<Form.Item label="平台奖励" name="platformBonus">
					<Input />
				</Form.Item>
				<Form.Item label="奖池金额" name="poolAmount">
					<Input />
				</Form.Item>
			</Form>
		</Modal>
	);
});

FormModal.displayName = "FormModal";

export default FormModal;
