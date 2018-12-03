interface CountResult {
    user: number;
    registry: number;
    model: number;
    updatedAt: string;
}

interface PojoResult {
    user: number;
    registry: number;
    model: number;
    createdAt: string;
    updatedAt: string;
    objectId: string;
}

export {
    CountResult,
    PojoResult
};