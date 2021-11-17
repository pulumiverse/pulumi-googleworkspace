import * as pulumi from "@pulumi/pulumi";
import * as googleWorkspace from "@pulumi/googleworkspace";

const group = new googleWorkspace.group.Group("pulumi-sample", {
    email: 'pulumi-sample@foo.com'
});

export const groupName = group.name;