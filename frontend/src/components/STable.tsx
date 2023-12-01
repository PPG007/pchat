import { Table, TablePaginationConfig } from "antd";
import { SorterResult } from "antd/es/table/interface";
import { FC, useEffect, useState } from "react";
import { ListCondition } from "../pb/common/request.ts";
import { ListResponse, STableProps } from "./props.ts";

const STable = <T = any>(): FC<STableProps<T>> => {
  return (props) => {
    const [isLoading, setIsLoading] = useState(false);
    const [data, setData] = useState<ListResponse<any>>({
      items: [],
      total: 0,
    });
    const fetchData = async (listCondition: ListCondition) => {
      setIsLoading(true);
      try {
        setData(await props.data(listCondition));
      } finally {
        setIsLoading(false);
      }
    }
    useEffect(() => {
      fetchData({
        page: 1,
        perPage: 10,
        orderBy: [],
      })
    }, [])
    const pageConfig: TablePaginationConfig = {
      ...props.pagination,
      total: data?.total,
      showSizeChanger: true,
    }
    const formatSorter = (sorter: SorterResult<any> | SorterResult<any>[]): Array<string> => {
      const orderBy: Array<string> = [];
      // TODO:
      return orderBy;
    }
    return (
      <Table
        {...props}
        onChange={({current, pageSize}, _filter, sorter) => {
          const listCondition: ListCondition = {
            page: current || 1,
            perPage: pageSize || 10,
            orderBy: formatSorter(sorter),
          }
          fetchData(listCondition);
        }}
        loading={isLoading}
        dataSource={data?.items}
        pagination={pageConfig}
      />
    )
  }
}

export default STable;