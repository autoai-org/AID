import { useState } from 'react'
import { Link } from 'react-router-dom'
import { SearchIcon } from '@heroicons/react/solid'

export default function Filters(props: any) {
    const [keyword, setKeyword] = useState("")
    function handleKWChange(e: any) {
        setKeyword(e.target.value)
    }
    function handleSearch() {
        props.search(true)
    }
    return (
        <div className="flex" >
            <div className="absolute rounded-md mt-1 flex shadow-sm" >
                <div className="relative flex items-stretch flex-grow focus-within:z-10">
                    <input
                        type="text"
                        id="keywords"
                        value={keyword}
                        onChange={handleKWChange}
                        className="focus:ring-blue-500 focus:border-blue-500 block w-full rounded-none rounded-l-md sm:text-sm border-gray-300"
                        placeholder="Search..."
                    />
                </div>
                <Link to={`/search?kw=${keyword}`}>
                    <button
                        type="button"
                        onClick={handleSearch}
                        className="-ml-px relative inline-flex items-center space-x-2 px-4 py-2 border border-gray-300 text-sm font-medium rounded-r-md bg-gray-50 hover:bg-gray-100 focus:ring-1 focus:ring-blue-500 focus:border-blue-500"
                    >
                        <SearchIcon className="h-5 w-5 text-gray-400" aria-hidden="true" />
                    </button>
                </Link>

            </div>
            <p className="mt-20">
                No Filters Available.
                We are working on this, stay tuned!
            </p>
        </div>
    )
}