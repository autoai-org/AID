// Copyright (c) 2021 Xiaozhe Yao et al.
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import { useEffect, useState } from "react";
import DoubleColumn from '../Layouts/DoubleColumn'
import MainLayout from "../Layouts/MainLayout";
import { parseQuery } from "../services/utility";
import { findModelsByKeyword, analyseRepos } from "../services/api"
import ModelList from "../components/Models/ModelList";

export default function SearchResult() {
    const [packages, setPackages] = useState([])
    const [loading, setLoading] = useState(false)

    useEffect(() => {
        let keyword = parseQuery("kw")
        console.log(keyword)
        setLoading(true)
        findModelsByKeyword(keyword).then(function (res: any) {
            console.log(res.data)
            analyseRepos(res.data).then(function (res) {
                setPackages(res)
                setLoading(false)
            });
        })
    }, [])
    return (
        <MainLayout>
            {loading &&
                <p>is loading...</p>
            }
            {!loading &&
                <DoubleColumn main={<ModelList models={packages} />}></DoubleColumn>
            }

        </MainLayout>
    )
}