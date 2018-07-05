namespace go addservice

struct AddRequest {
    1: i64 a;
    2: i64 b;
}

struct AddResponse {
    1: i64 v;
    2: string err;
}

service AddService {
    AddResponse add(1: AddRequest request);
}
