cfn-show
========

**Shows one Parameter, Mapping, Output, or Resource from a CloudFormation template.**

This tool is good for digging into one item in a huge CloudFormation template. Use it like so:

    cfn-show template-file item-type/item-name

So, to check out a Resource named WebInstance in a file called GreatWebsite.template:

    cfn-show GreatWebsite.template Resource/WebInstance
    
The item will be pretty printed to STDOUT.
