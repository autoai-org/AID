import ModelCard from './ModelCard'
import Divider from '../Common/Divider'
export default function ModelList(props: any) {
    return (
        <ul>
            {props.models.map((model: any, index: number) => (
            <li>
                <ModelCard></ModelCard>
                <Divider/>
            </li>))}
        </ul>
    )
}