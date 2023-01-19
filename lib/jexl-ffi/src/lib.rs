use jexl_eval::Evaluator;
use serde_json::{from_str, json, Value};
use std::ffi::{CStr, CString};

pub struct JexlEngine<'a> {
    evaluator: Evaluator<'a>,
    context: Value,
    script: String,
}

impl JexlEngine<'static> {
    fn new(context: Value, script: String) -> JexlEngine<'static> {
        JexlEngine {
            evaluator: Evaluator::new(),
            context,
            script,
        }
    }

    fn run(&self) -> String {
        let res = self
            .evaluator
            .eval_in_context(&self.script, self.context.clone())
            .expect("Failed to eval script");
        return serde_json::to_string(&res).expect("Failed to deserialize result");
    }
}

#[no_mangle]
pub extern "C" fn new_engine(
    context_ptr: *const libc::c_char,
    script_ptr: *const libc::c_char,
) -> *mut JexlEngine<'static> {
    let context_string = to_string(context_ptr);
    let context: Value = from_str(&context_string).unwrap();
    let script = to_string(script_ptr);
    Box::into_raw(Box::new(JexlEngine::new(context, script)))
}

#[no_mangle]
pub extern "C" fn free_engine(ptr: *mut JexlEngine) {
    if ptr.is_null() {
        return;
    }
    unsafe {
        Box::from_raw(ptr);
    }
}

#[no_mangle]
pub extern "C" fn run_engine(ptr: *mut JexlEngine<'static>) -> *const libc::c_char {
    let engine = unsafe {
        assert!(!ptr.is_null());
        &mut *ptr
    };

    let result = engine.run();
    CString::new(result).expect("Engine run failed").into_raw()
}

fn to_string(ptr: *const libc::c_char) -> String {
    let cstr = unsafe { CStr::from_ptr(ptr) };
    return cstr.to_str().unwrap().to_string();
}

#[no_mangle]
pub extern "C" fn eval(script: *const libc::c_char) -> *const libc::c_char {
    let message_cstr = unsafe { CStr::from_ptr(script) };
    let message = message_cstr.to_str().unwrap();

    let evaluator = Evaluator::new();
    let eval_value = evaluator.eval(message).unwrap();
    let val_str = serde_json::to_string(&eval_value).unwrap();
    return CString::new(val_str).unwrap().into_raw();
}

#[cfg(test)]
pub mod test {

    use super::*;
    use serde_json::json;
    use std::ffi::CString;

    #[test]
    fn oneshot_eval() {
        let expr = CString::new("6 * 12 + 5 / 2.6").unwrap().into_raw();
        let ffi_val = eval(expr);
        let val = unsafe { CStr::from_ptr(ffi_val) };
        assert_eq!(val.to_str().unwrap(), "73.92307692307692")
    }

    #[test]
    fn engine() {
        let context = CString::new(json!({"a": {"b": 2.0}}).to_string())
            .unwrap()
            .into_raw();
        let script = CString::new("a.b").unwrap().into_raw();
        let engine = new_engine(context, script);
        let res_raw = run_engine(engine);
        let res = unsafe { CStr::from_ptr(res_raw) };
        assert_eq!(res.to_str().unwrap(), "2.0");
        free_engine(engine);
    }
}
