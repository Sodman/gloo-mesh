---
title: "meshctl CLI"
menuTitle: meshctl CLI
description: Installing the meshctl CLI
weight: 120
---

Use the Gloo Mesh command line interface (CLI) tool, `meshctl`, to set up Gloo Mesh, register clusters, describe your Gloo Mesh resources, and more.

## Quick installation

Install the latest version of `meshctl`. Make sure to add `meshctl` to your PATH so you can run commands like `meshctl version`.

```shell
curl -sL https://run.solo.io/meshctl/install | sh
```

## Installing a specific version of the CLI

1.  Go the Gloo Mesh releases GitHub page.
    {{< tabs >}}
    {{< tab name="Enterprise">}}
    https://github.com/solo-io/gloo-mesh-enterprise/releases
    {{< /tab >}}
    {{< tab name="Open Source">}}
    https://github.com/solo-io/gloo-mesh/releases
    {{< /tab >}}
    {{< /tabs >}}
2.  Click the version that you want to install and, from the **Assets** section, download the `meshctl` package for your operating system.
3.  If you use macOS or Linux, complete the following steps.

    1.  Move the downloaded file to the `/usr/local/bin` directory on your system.
        ```shell
        mv ~/Downloads/meshctl-darwin-amd64 ~/.gloo-mesh/bin/meshctl
        ```
    2.  Make the downloaded file executable.
        ```shell
        chmod +x /usr/local/bin/meshctl
        ```
4.  Add `meshctl` on your PATH system variable for global access on the command line. The steps vary depending on your operating system. The following eample is for macOS. For more information, see ([Windows](https://helpdeskgeek.com/windows-10/add-windows-path-environment-variable/), [macOS](https://osxdaily.com/2014/08/14/add-new-path-to-path-command-line/), or [Linux](https://linuxize.com/post/how-to-add-directory-to-path-in-linux/).
    1.  Add the `~/.gloo-mesh/bin` to your PATH.
        ```shell
        export PATH=$HOME/.gloo-mesh/bin:$PATH
        ```
    2.  Verify that the PATH is updated.
        ```shell
        echo $PATH
        ```    
        Example output:
        ```
        /usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin
        ```
5.  Verify that you can run `meshctl` commands.
    ```shell
    meshctl version
    ```

    {{% notice tip %}}
    On macOS, you might see the following warning: `“meshctl” cannot be opened because it is from an unidentified developer.` In the Finder app, navigate to the `~/.gloo-mesh/bin/meshctl` executable file, right-click the file, click **Open**, and confirm that you want to open the file. For more information, try searching the warning and following a guide such as [this blog](https://www.howtogeek.com/205393/gatekeeper-101-why-your-mac-only-allows-apple-approved-software-by-default/).
    {{% /notice %}}

Good job! You now have the version of `meshctl` that you want installed.

Do you have multiple cluster environments that require different versions of Gloo Mesh, Istio, and Kubernetes? Consider downloading each `meshctl`, `istioctl`, and `kubectl` version binary file to a separate directory. Then, you can set up an alias in your local command line interface profile to point to the binary file directory that matches the version of the cluster environment that you want to work with.

## Updating the CLI

Updating the `meshctl` CLI does _not_ [upgrade the Gloo Mesh version]({{% versioned_link_path fromRoot="/operations/upgrading/" %}}) that you run in your clusters.

1. [Uninstall the your current CLI version]({{<ref "#uninstalling-the-cli>}}).
2. Install the CLI. You can install the [latest]({{<ref "#quick-installation>}})) or a [specific version]({{<ref "#installing-a-specific-version-of-the-cli>}}).

## Uninstalling the CLI

To uninstall `meshctl`, you can delete the executable file from your system, such as on macOS in the following example.

```shell
rm ~/.gloo-mesh/bin/meshctl
```
