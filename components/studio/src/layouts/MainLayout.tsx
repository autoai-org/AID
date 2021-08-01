import Body from './Body'
import Loading from '../components/Common/Loading'
export default function MainLayout() {
    return (
        <div className="mx-auto">
            <main>
                <Loading />
                <Body/>
            </main>
        </div>
    )
}