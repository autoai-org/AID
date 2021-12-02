# coding: utf-8
from mlpm.app import aidserver

prediction_schema = aidserver.create_validator(
    "predict_request",
    {
        "type": "object",
        "properties": {
            "input": {"type": "string"},
            "file": {
                "type": "string",
                "format": "binary",
            },
        },
    },
)
