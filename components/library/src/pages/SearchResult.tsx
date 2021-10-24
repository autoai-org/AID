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
import Filters from '../components/Models/ModelFilters'
export default function SearchResult() {
    const [packages, setPackages] = useState([])
    const [loading, setLoading] = useState(false)
    const [progress, setProgress] = useState(0)
    useEffect(() => {
        let keyword = parseQuery("kw")
        setLoading(true)
        setProgress(1)
        findModelsByKeyword(keyword).then(function (res: any) {
            setProgress(2)
            analyseRepos(res.data, false).then(function (res) {
                setPackages(res)
                setLoading(false)
            });
        })
    }, [])
    return (
        <MainLayout>
            {loading && progress === 1 &&
                <div style={{ height: '100vh' }}>
                    <div className="pt-36" style={{ marginLeft: 'auto', marginRight: 'auto' }}>
                        <svg className="animate-spin h-24 w-24 mr-3 ..." viewBox="0 0 24 24" style={{ marginLeft: 'auto', marginRight: 'auto' }}>
                            <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                            <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                        </svg>
                        <p className="text-2xl" style={{ marginLeft: 'auto', marginRight: 'auto', textAlign: 'center' }}>Listing Repositories...</p>
                        <p className="text-2xl" style={{ marginLeft: 'auto', marginRight: 'auto', textAlign: 'center' }}>1/2 Please wait patiently, it may take few minutes...</p>
                    </div>
                </div>
            }
            {loading && progress === 2 &&
                <div style={{ height: '100vh' }}>
                    <div className="pt-36" style={{ marginLeft: 'auto', marginRight: 'auto' }}>
                        <svg className="animate-spin h-24 w-24 mr-3 ..." viewBox="0 0 24 24" style={{ marginLeft: 'auto', marginRight: 'auto' }}>
                            <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                            <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                        </svg>
                        <p className="text-2xl" style={{ marginLeft: 'auto', marginRight: 'auto', textAlign: 'center' }}>Analysing and Fetching Metadata...</p>
                        <p className="text-2xl" style={{ marginLeft: 'auto', marginRight: 'auto', textAlign: 'center' }}>2/2 Please wait patiently, it may take few minutes...</p>
                    </div>
                </div>
            }
            {!loading &&
                <DoubleColumn left={<Filters/>} main={<ModelList models={packages}/>}></DoubleColumn>
            }
        </MainLayout>
    )
}