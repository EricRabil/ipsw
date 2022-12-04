#!/bin/bash

echo "i hope you know that this is a house of cards"

IPSW=./ipsw

if [ -z "$WORK_DIR" ]; then
    WORK_DIR="/tmp/kdeek"
fi

function ensure_volume_group {
    if [ -z "$VOLUME_GROUP" ]; then
        read -p "volume group: " VOLUME_GROUP
    fi
}

function ensure_kernel_path {
    if [ -z "$KERNEL_PATH" ]; then
        read -p "kernel path:" KERNEL_PATH;
    fi
}

function ensure_kdk_path {
    if [ -z "$KDK_OUTPUT_PATH" ]; then
        read -p "kernel path:" KDK_OUTPUT_PATH;
    fi
}

function decompress_kernelcache {
    ensure_volume_group

    PREBOOT_PATH="/System/Volumes/Preboot/$VOLUME_GROUP"
    ACTIVE_BOOT=$(cat "$PREBOOT_PATH/boot/active")
    COMPRESSED_KERNELCACHE_PATH="$PREBOOT_PATH/boot/$ACTIVE_BOOT/System/Library/Caches/com.apple.kernelcaches/kernelcache"

    $IPSW kernel dec -V "$COMPRESSED_KERNELCACHE_PATH" --km --output "$WORK_DIR"
}

function extract_kexts {
    ensure_volume_group

    PREBOOT_PATH="/System/Volumes/Preboot/$VOLUME_GROUP"
    ACTIVE_BOOT=$(cat "$PREBOOT_PATH/boot/active")
    COMPRESSED_KERNELCACHE_PATH="$PREBOOT_PATH/boot/$ACTIVE_BOOT/System/Library/Caches/com.apple.kernelcaches/kernelcache"
    DECOMPRESSED_KERNELCACHE_PATH="$WORK_DIR$COMPRESSED_KERNELCACHE_PATH.decompressed"

    $IPSW -V macho info --output "$WORK_DIR" -z -x "$DECOMPRESSED_KERNELCACHE_PATH"
}

function assemble_kdk {
    ensure_kdk_path

    echo "mkdir -p $KDK_OUTPUT_PATH"
    sudo mkdir -p "$KDK_OUTPUT_PATH/System/Library/Extensions"

    ensure_kernel_path

    sudo rsync -rav /System/Library/Extensions "$KDK_OUTPUT_PATH/System/Library"
    sudo rsync -rav "$WORK_DIR/System/Library/Extensions" "$KDK_OUTPUT_PATH/System/Library"
    sudo rsync -rav /System/Library/Kernels "$KDK_OUTPUT_PATH/System/Library"

    PLUGINS="$KDK_OUTPUT_PATH/System/Library/Extensions/System.kext/PlugIns"

    ./symbolsets.sh allsymbols "$KERNEL_PATH" "$WORK_DIR/allsymbols"
    ./symbolsets.sh extract "$KERNEL_PATH"
    sudo ./symbolsets.sh proxy "$WORK_DIR/allsymbols" com.apple.kpi.bsd "$PLUGINS/BSDKernel.kext/BSDKernel"
    sudo ./symbolsets.sh proxy "$WORK_DIR/allsymbols" com.apple.kpi.unsupported "$PLUGINS/Unsupported.kext/Unsupported"
    sudo ./symbolsets.sh proxy "$WORK_DIR/allsymbols" com.apple.kpi.private "$PLUGINS/Private.kext/Private"
    sudo ./symbolsets.sh proxy "$WORK_DIR/allsymbols" com.apple.kpi.mach "$PLUGINS/Mach.kext/Mach"
    sudo ./symbolsets.sh proxy "$WORK_DIR/allsymbols" com.apple.kpi.dsep "$PLUGINS/MACFramework.kext/MACFramework"
    sudo ./symbolsets.sh proxy "$WORK_DIR/allsymbols" com.apple.kpi.libkern "$PLUGINS/Libkern.kext/Libkern"
    sudo ./symbolsets.sh proxy "$WORK_DIR/allsymbols" com.apple.kpi.kcov "$PLUGINS/Kcov.kext/Kcov"
    sudo ./symbolsets.sh proxy "$WORK_DIR/allsymbols" com.apple.kpi.kasan "$PLUGINS/Kasan.kext/Kasan"
    sudo ./symbolsets.sh proxy "$WORK_DIR/allsymbols" com.apple.kpi.iokit "$PLUGINS/IOKit.kext/IOKit"
}

function make_kernelcache {
    ensure_kdk_path
    ensure_kernel_path

    kmutil create \
        --allow-missing-kdk \
        --kdk "$KDK_OUTPUT_PATH" \
        -a arm64e \
        -z \
        -V release \
        -n boot \
        -B rel.kc \
        -k "$KERNEL_PATH" \
        -r /System/Library/Extensions \
        -r /System/Library/DriverExtensions \
        -r "$KDK_OUTPUT_PATH/System/Library/Extensions" \
        -r "$KDK_OUTPUT_PATH/System/Library/Extensions/System.kext/PlugIns" \
        -x $(echo $KDK_OUTPUT_PATH/System/Library/Extensions/*.kext $KDK_OUTPUT_PATH/System/Library/Extensions/System.kext/PlugIns/*.kext | tr ' ' '\n' | awk '{print " -b "$1; }' | tr '\n' ' ')
}

case $1 in
    "prep-kernelcache")
        decompress_kernelcache
        ;;
    "extract-kexts")
        extract_kexts
        ;;
    "assemble-kdk")
        assemble_kdk
        ;;
    "make-kernelcache")
        make_kernelcache
        ;;
    "all")
        decompress_kernelcache extract_kexts assemble_kdk make_kernelcache
        ;;
    *)
        echo "wat"
        ;;
esac
