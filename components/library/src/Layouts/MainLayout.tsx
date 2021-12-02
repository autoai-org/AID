// Copyright (c) 2021 Xiaozhe Yao et al.
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import AIDHeader from "../components/Header";
import AIDFooter from "../components/Footer";

export default function MainLayout(props: any) {
  return (
    <div className="flex flex-col h-screen">
      <AIDHeader />
      <main className="flex-grow">
        <div>{props.children}</div>
      </main>
      <AIDFooter />
    </div>
  );
}
