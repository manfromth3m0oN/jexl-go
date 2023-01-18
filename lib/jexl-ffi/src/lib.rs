use jexl_eval::Evaluator;
use std::ffi::{CStr, CString};

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
    use std::ffi::CString;

    #[test]
    fn simulated_main_function() {
        let expr = CString::new("6 * 12 + 5 / 2.6").unwrap().into_raw();
        let ffi_val = eval(expr);
        let val = unsafe { CStr::from_ptr(ffi_val) };
        println!("val: {}", val.to_str().unwrap());
        assert_eq!(val.to_str().unwrap(), "73.92307692307692")
    }
}
