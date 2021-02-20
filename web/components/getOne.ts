import { GetServerSidePropsContext, GetServerSidePropsResult } from "next";
import * as grpc from "@grpc/grpc-js";
import * as jspb from 'google-protobuf'
import { Error, StatusCode } from "grpc-web";

interface Instance extends jspb.Message {
  getId: () => number;
  toObject: () => any;
}

interface getByIDRequest<T extends Instance> extends jspb.Message {
  setId: (id: number) => getByIDRequest<T>;
}

interface getByIDResponse<T extends Instance> extends jspb.Message {
  getResult: () => T;
}

interface getByIDClient<T extends Instance> {
  getOneByID(request: getByIDRequest<T>, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: getByIDResponse<T>) => void): grpc.ClientUnaryCall;
}

interface pageProp<T extends Instance> {
  id: number;
  result?: T;
  errorCode?: number;
}

type getServerSideFunc<T extends Instance> = (context: GetServerSidePropsContext) => Promise<GetServerSidePropsResult<pageProp<T>>>;

const httpStatusCode = (grpcCode: number): number => {
  switch (grpcCode) {
    case StatusCode.OK:
      return 200
    case StatusCode.NOT_FOUND:
      return 404
    case StatusCode.DEADLINE_EXCEEDED:
      return 503
    default:
      return 500
  }
};

export const getOneByID = async<T extends Instance>(
  context: GetServerSidePropsContext,
  request: getByIDRequest<T>,
  client: getByIDClient<T>,
  authorizationToken?: string,
): Promise<GetServerSidePropsResult<pageProp<T>>> => {
  let id: number;

  if (Array.isArray(context.params.id)) {
    id = parseInt(context.params.id[0], 10);
  } else {
    id = parseInt(context.params.id, 10);
  }

  const props: pageProp<T> = {
    id,
    result: null,
    errorCode: null,
  };

  const metadata = new grpc.Metadata();

  if (authorizationToken) {
    metadata.set('Authorization', 'Bearer ' + authorizationToken);
  }

  request.setId(id);
  const p = new Promise((resolve, reject) =>
    client.getOneByID(request, metadata, (err: Error, response: getByIDResponse<T>) => {
      if (err) {
        return reject(err);
      }
      return resolve(response);
    })
  );

  // await (p);
  await p
    .then(
      (response: getByIDResponse<T>) => {
        props.result = response.getResult().toObject();
      },
      (e: Error) => {
        props.errorCode = e.code;
      }
    )
    .catch(() => {
      props.errorCode = StatusCode.UNKNOWN;
    });

  if (props.errorCode) {
    context.res.statusCode = httpStatusCode(props.errorCode)
  }

  return {
    props,
  };
};