import ModelCard from './ModelCard'
export default function ModelList(props: any) {
    return (
        <ul>
            {props.models.map((model: any, index: number) => (
            <li>
                <ModelCard model={model}></ModelCard>
            </li>))}
        </ul>
    )
}