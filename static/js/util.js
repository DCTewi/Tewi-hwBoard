function checkbyid(formname, id, pattern) {
    var x = document.forms[formname][id].value
    console.log(x)
    if (x == null || x == "") {
        $("#" + id).addClass("is-invalid")
        return false
    }

    var flag = x.search(pattern)
    var ad = "is-invalid",  rm = "is-valid"
    if (flag == 0) {
        ad = [rm, rm = ad][0]
    }
    $("#" + id).addClass(ad)
    $("#" + id).removeClass(rm)

    return (flag == 0)
}
