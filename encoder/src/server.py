import time
import grpc  # type: ignore
from concurrent import futures
import torch
from loguru import logger
from sentence_transformers import SentenceTransformer  # type: ignore
import text_encoder_pb2  # type: ignore
import text_encoder_pb2_grpc  # type: ignore
import os

logger.info("Cuda available: {}".format(torch.cuda.is_available()))

_model = SentenceTransformer("sentence-transformers/LaBSE")


class TextEncoderService(text_encoder_pb2_grpc.TextEncoderServiceServicer):

    def EncodeText(self, request, context):
        text = request.text
        vector = get_encoded_text(text)
        # Convert the torch tensor to a list to send as response
        vector_list = vector.tolist()
        return text_encoder_pb2.VectorResponse(vector=vector_list)


def get_encoded_text(text: str) -> torch.Tensor:
    start = time.time()
    vector = _model.encode(text, convert_to_tensor=True)
    end = time.time()
    logger.info("Inference time: {}".format(end - start))
    return vector.squeeze()


def serve():
    cpu_count = os.cpu_count()
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=cpu_count))
    text_encoder_pb2_grpc.add_TextEncoderServiceServicer_to_server(
        TextEncoderService(), server
    )
    server.add_insecure_port("[::]:50051")
    server.start()
    server.wait_for_termination()


if __name__ == "__main__":
    serve()
