// Copyright (c) 2021 Xiaozhe Yao et al.
// 
// This software is released under the MIT License.

// https://opensource.org/licenses/MIT
export default function DoubleColumn(props: any) {
  return (
    <>
      <div className="flex-grow w-full max-w-7xl mx-auto lg:flex" style={{height:'100vh'}}>
        <div className="flex-1 min-w-0 xl:flex">
          <div className="xl:border-b-0 xl:flex-shrink-0 xl:w-64 xl:border-r xl:border-gray-200">
            <div className="h-full pl-4 pr-6 py-6 sm:pl-6 lg:pl-8 xl:pl-0">
              {/* Start left column area */}
              <div className="h-full relative" style={{ minHeight: '12rem' }}>
                <div className="absolute inset-0  rounded-lg" />
              </div>
              {/* End left column area */}
            </div>
          </div>

          <div className="lg:min-w-0 lg:flex-1">
            <div className="h-full py-6 px-4 sm:px-6 lg:px-8">
              {/* Start main area*/}
              {props.main}
              {/* End main area */}
            </div>
          </div>
        </div>
      </div>
    </>
  )
}