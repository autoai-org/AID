import { useState, useEffect } from "react"
import { analyseRepos, findModelsByKeyword } from "../../services/api"
import { getInitials } from "../../services/utility";
export default function Leaderboard() {
    const [packages, setPackages] = useState([]);
    const [vendors, setVendors] = useState([]);

    useEffect(() => {
        findModelsByKeyword("").then(function (res: any) {
            analyseRepos(res.data, true).then(function (res: any) {
                setPackages(res);
                let vendors_count = res.reduce(function(allVendors: any, repo: any) {
                    if (repo.vendor in allVendors) {
                        allVendors[repo.vendor] ++;
                    } else {
                        allVendors[repo.vendor] = 1;
                    }
                    return allVendors;
                }, {})
                setVendors(vendors_count)
            });
        })
    }, [])
    return (
        <div className="bg-white">
            <div className="max-w-7xl mx-auto py-12 px-4 text-center sm:px-6 lg:px-8 lg:py-24">
                <div className="space-y-8 sm:space-y-12">
                    <div className="space-y-5 sm:mx-auto sm:max-w-xl sm:space-y-4 lg:max-w-5xl">
                        <h2 className="text-3xl font-extrabold tracking-tight sm:text-4xl">Leaderboard</h2>
                        <p className="text-xl text-gray-500">
                            The most active contributors on the platform
                        </p>
                    </div>
                    <h3 className="text-2xl tracking-tight sm:text-4xl">Vendors</h3>
                    <ul
                        className="mx-auto grid grid-cols-2 gap-x-4 gap-y-8 sm:grid-cols-4 md:gap-x-6 lg:max-w-5xl lg:gap-x-8 lg:gap-y-12 xl:grid-cols-6"
                    >
                        {Object.keys(vendors).map((keyName:any, keyIdx) => (
                            <li key={keyName}>
                                <div className="space-y-4">
                                <span className="inline-flex items-center justify-center h-24 w-24 rounded-full bg-gray-500">
                                        <span className="text-4xl leading-none text-white">{getInitials(keyName)}</span>
                                    </span>
                                    <div className="space-y-2">
                                        <div className="text-xs font-medium lg:text-sm">
                                            <h3>{keyName}</h3>
                                            <p className="text-indigo-600"><a href={'https://github.com/'+keyName}>GitHub</a></p>
                                        </div>
                                    </div>
                                </div>
                            </li>
                        ))}
                    </ul>
                    <h3 className="text-2xl tracking-tight sm:text-4xl">Packages</h3>
                    <ul
                        className="mx-auto grid grid-cols-2 gap-x-4 gap-y-8 sm:grid-cols-4 md:gap-x-6 lg:max-w-5xl lg:gap-x-8 lg:gap-y-12 xl:grid-cols-6"
                    >
                        {packages.map((each: any, idx: any) => (
                            <li key={idx}>
                                <div className="space-y-4">
                                <img className="mx-auto h-20 w-20 rounded-full lg:w-24 lg:h-24" src={each.avatar} alt=""/>
                                    <div className="space-y-2">
                                        <div className="text-xs font-medium lg:text-sm">
                                            <h3>{each.name}</h3>
                                            <p className="text-indigo-600"><a href={each.githubURL} >GitHub</a></p>
                                        </div>
                                    </div>
                                </div>
                            </li>
                        ))}
                    </ul>
                </div>
            </div>
        </div>
    )
}
