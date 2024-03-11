package main.java.Api;

import common.java.Apps.AppContext;
import common.java.String.StringHelper;

public class Role {
    public String getRole() {
        var allRole = AppContext.current().roles().all();
        return StringHelper.join(allRole);
    }
}
