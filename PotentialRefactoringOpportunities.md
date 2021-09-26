# Potential Refactoring Opportunities
## Overview
This document serves as both a reflection on what i've built as well as what i think could be improved. Since this is a small project, i wanted to keep it as simple as possible without going overly complicated, and as such there are things i would have updated. I'll break this reflection down into front and back end, with a general notes section for things that apply to both.

## General Notes
The biggest thing to include here is the duplication of generated Protobuf files. I initially had these in their own package called searchtermsprotobuf, but when building the docker images this wasn't possible. To remedy this, i've copies the generated protobuf files into each project, however it means we have duplicated code and as the project grows it could lead to more things to update. This isn't necissarily bad, as if you wanted to add more services, there wouldn't be any tight coupling, but it just mean's more work. It's all trade offs.

This leads to my second point, which is i could also have each service in it's own repo. This would allow us to turn each project into a module a whole lot easier, and it wouldn't have just a singular name. But this comes down to a mono repo/multi repo debate, and i'm always happy to go with the company culture on what they use as long as it's doing it's job.

## Front End
For the most part I'm pretty happy with this. The front end acts as a gateway, and takes in a request HTTP and makes a gRPC call, returning the result. Because of this however i didn't include any tests, since there really wasn't any logic. This may benefit from a different kind of testing.

## Backend
This is pretty solid i think. I'm happy with the tests, though these could always be refined since I'm still new to the table testing method Go uses. However the big talking point on this is the use of a map. To ensure that this was as portable as possible, i just used a map as the backing store. This could be changed easily, it would just need some extra code. But right now it's not surrounded by any kind of lock or threading mechanism. So this could turn out to have some race conditions. This should be revisited.