// Copyright (c) 2021 Xiaozhe Yao et al.
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import { DataGrid, GridColDef, GridValueGetterParams } from '@material-ui/data-grid';
import { useQuery } from '@apollo/client';
import { ALL_REPOSITORIES } from '../../services/apis/queries'

const columns: GridColDef[] = [
    { field: 'uid', headerName: 'UID', width: 150 },
    { field: 'name', headerName: 'Name', width: 200 },
    { field: 'status', headerName: 'Status', width: 200 },
  ];

export default function RepoTable() {
    const { loading, error, data } = useQuery(ALL_REPOSITORIES);    
    if (loading) return <p>Loading...</p>;
    if (error) return <p>Error :(</p>;
    const rows = data.repositories;
    console.log(rows)
    return (
        <div style={{ height: 400, width: '100%' }}>
            <DataGrid rows={data.repositories} columns={columns} pageSize={5} checkboxSelection />
        </div>
        
    )
}