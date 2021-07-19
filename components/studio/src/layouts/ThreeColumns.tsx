// Copyright (c) 2021 Xiaozhe Yao et al.
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

export default function ThreeColumnsLayout(props: any) {
    return (
        <>
            {/* Background color split screen for large screens */}
            <div className="relative min-h-screen flex flex-col">

                {/* 3 column wrapper */}
                <div className="flex-grow w-full max-w-7xl mx-auto xl:px-8 lg:flex">
                    {/* Left sidebar & main wrapper */}
                    <div className="flex-1 min-w-0 xl:flex">
                        <div className="border-b border-gray-200 xl:border-b-0 xl:flex-shrink-0 xl:w-64 xl:border-r xl:border-gray-200">
                            <div className="h-full pl-4 pr-6 py-6 sm:pl-6 lg:pl-8 xl:pl-0">
                                {/* Start left column area */}
                                <div className="h-full relative" style={{ minHeight: '12rem' }}>
                                    {props.left}
                                </div>
                                {/* End left column area */}
                            </div>
                        </div>

                        <div className="lg:min-w-0 lg:flex-1">
                            <div className="h-full py-6 px-4 sm:px-6 lg:px-8">
                                {/* Start main area*/}
                                <div className="relative h-full" style={{ minHeight: '36rem' }}>
                                    {props.middle}
                                </div>
                                {/* End main area */}
                            </div>
                        </div>
                    </div>

                    <div className="pr-4 sm:pr-6 lg:pr-8 lg:flex-shrink-0 lg:border-l lg:border-gray-200 xl:pr-0">
                        <div className="h-full pl-6 py-6 lg:w-80">
                            {/* Start right column area */}
                            <div className="h-full relative" style={{ minHeight: '16rem' }}>
                                {props.right}
                            </div>
                            {/* End right column area */}
                        </div>
                    </div>
                </div>
            </div>
        </>
    )
}