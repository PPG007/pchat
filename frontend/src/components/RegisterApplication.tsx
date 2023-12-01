import { ColumnType } from "antd/es/table";
import { FC } from "react";
import i18n from "../../i18n";
import { User } from "../engine";
import { RegisterApplicationDetail } from "../pb/user/request.ts";
import STable from "./STable.tsx";

const columns: Array<ColumnType<RegisterApplicationDetail>> = [
  {
    title: i18n.t('registerApplication.email'),
    dataIndex: 'email',
  },
  {
    title: i18n.t('registerApplication.status'),
    dataIndex: 'status',
  },
  {
    title: i18n.t('registerApplication.reason'),
    dataIndex: 'reason',
  },
  {
    title: i18n.t('registerApplication.rejectReason'),
    dataIndex: 'rejectReason',
  },
  {
    title: i18n.t('registerApplication.createdAt'),
    dataIndex: 'createdAt',
    sorter: true,
  }
];

const RegisterApplication: FC = () => {
  const Table = STable<RegisterApplicationDetail>();
  return (
    <>
      <Table
        data={
          async (listCondition) => {
            const resp = await User.fetchRegisterApplications({
              status: [],
              listCondition,
            })
            return {
              total: resp.data.total,
              items: resp.data.items,
            }
          }
        }
        columns={columns}
        rowKey={'id'}
      />
    </>
  )
}

export default RegisterApplication;