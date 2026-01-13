import { ExclamationCircleOutlined, PlusCircleOutlined, CloudDownloadOutlined, EditOutlined, DeleteOutlined } from "@ant-design/icons";
import { ProColumns, ProFormInstance, ProTable } from "@ant-design/pro-components";
import { Modal, Space } from "antd";
import React, { useRef } from "react";
import HocAuth from "@/components/HocAuth";
import LoadingButton from "@/components/LoadingButton";
import { pagination } from "@/config/proTable";
import { formatDataForProTable, saveExcelBlob } from "@/utils";
import { message, modal } from "@/hooks/useMessage";
import {
  ChallengeCheckinImageModel,
  getCheckinImagePageApi,
  exportCheckinImageApi,
  addCheckinImageApi,
  updateCheckinImageApi,
  delCheckinImageApi,
} from "@/api/app/challenge/checkinImage";
import type { ActionType } from "@ant-design/pro-components";
import FormModal from "./components/FormModal";
import type { FormModalRef } from "./components/FormModal";

const CheckinImage: React.FC = () => {
  const actionRef = useRef<ActionType>();
  const formRef = useRef<ProFormInstance>();
  const formModalRef = useRef<FormModalRef>(null);

  const handleDelete = (id: number) => {
    Modal.confirm({
      title: "确认删除该打卡图片？",
      icon: <ExclamationCircleOutlined />,
      onOk: async () => {
        await delCheckinImageApi(id);
        message.success("删除成功");
        actionRef.current?.reload();
      },
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
          saveExcelBlob("challenge_checkin_image", await exportCheckinImageApi(formRef.current?.getFieldsValue()));
        } catch (err) {
          message.error("导出失败");
        } finally {
          done();
        }
      },
    });
  };

  const handleShowCreate = () => {
    formModalRef.current?.showAddFormModal();
  };

  const handleShowEdit = (record: ChallengeCheckinImageModel) => {
    if (!record.id) return;
    formModalRef.current?.showEditFormModal(record.id, record);
  };

  const columns: ProColumns<ChallengeCheckinImageModel>[] = [
    { title: "ID", dataIndex: "id", width: 80 },
    { title: "打卡ID", dataIndex: "checkinId", width: 120 },
    { title: "用户ID", dataIndex: "userId", width: 120 },
    { title: "图片URL", dataIndex: "imageUrl", width: 200, ellipsis: true },
    { title: "图片Hash", dataIndex: "imageHash", width: 180, hideInSearch: true },
    { title: "排序", dataIndex: "sortNo", width: 80, hideInSearch: true },
    { title: "状态", dataIndex: "status", width: 80 },
    { title: "上传时间", dataIndex: "createdAt", valueType: "dateTime", width: 180, hideInSearch: true },
    {
      title: "操作",
      valueType: "option",
      width: 200,
      fixed: "right",
      render: (_, record) => (
        <Space>
          <HocAuth permission={["app:challenge:checkinImage:edit"]}>
            <LoadingButton key="edit" type="link" size="small" icon={<EditOutlined />} onClick={() => handleShowEdit(record)}>
              编辑
            </LoadingButton>
          </HocAuth>
          <HocAuth permission={["app:challenge:checkinImage:delete"]}>
            <LoadingButton key="del" danger type="link" size="small" icon={<DeleteOutlined />} onClick={() => handleDelete(record.id!)}>
              删除
            </LoadingButton>
          </HocAuth>
        </Space>
      ),
    },
  ];

  const toolBarRender = () => [
    <HocAuth key="add" permission={["app:challenge:checkinImage:add"]}>
      <LoadingButton type="primary" key="addBtn" icon={<PlusCircleOutlined />} onClick={() => handleShowCreate()}>
        新增
      </LoadingButton>
    </HocAuth>,
    <HocAuth key="export" permission={["app:challenge:checkinImage:export"]}>
      <LoadingButton type="primary" key="exportBtn" icon={<CloudDownloadOutlined />} onClick={done => handleExport(done)}>
        导出
      </LoadingButton>
    </HocAuth>,
  ];

  return (
    <>
      <ProTable<ChallengeCheckinImageModel>
        rowKey="id"
        columns={columns}
        actionRef={actionRef}
        formRef={formRef}
        pagination={pagination}
        scroll={{ x: 1200 }}
        request={async params => {
          const { data } = await getCheckinImagePageApi(params);
          return formatDataForProTable<ChallengeCheckinImageModel>(data);
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

export default CheckinImage;
