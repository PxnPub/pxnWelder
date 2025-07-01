pxnWelder - "It welds the code together or something"








































pxnWelder - Automated Build Tool

I'm making a new build system and ci in golang. it will be modular and based on my current set of xbuild scripts. I want to use podman to build a project then delete the podman image, and if the latest commit has a tab then it will build a release rather than a snapshot.

I have a bash script, xBuild, which handles cleanup, workspace setup, configuration, building, packaging, and distribution of my software projects. It also handles git for me and has a --ci flag for use on a ci server, and it also detects if the latest commit has a tag then builds a full release, otherwise builds a snapshot. I am currently using this with jenkins.

I want to start a new project, pxnWelder, which will replace my old scripts and jenkins. weld will replace xBuild, and WelderCI will replace jenkins. I want to have a web frontend and a separate backend service to handle builds. It will use podman with the delete flag. Each new build will construct a new podman container, start up, run the weld --ci script, then deploy to a web server. It will also need to build rpm files.

I want to design a clean and modular system. Mostly golang, and some bash scripts are acceptable. It will also support skel files which are compared to files existing in a git workspace then update if different, and also generate other files from templates like pom.xml files.

weld - automated build tool
welder ci - ci service using podman and a website frontend service
welder badge - a web server for hosting badge svg generation
