import React from 'react'
import Table from '../Common/Table/Table.jsx'

export default function ModelList(props: any) {
    const columns = React.useMemo(() => [
        {
            Header: "Vendor",
            accessor: 'vendor',
        },
        {
            Header: "Name",
            accessor: 'name',
        },
    ], [])

    return (
        <div className="mt-2 flex flex-col">
            <Table columns={columns} data={props.models} />
        </div>
    )
}