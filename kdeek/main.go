package main

//#cgo CFLAGS: -x c -Wno-deprecated-declarations
// #cgo LDFLAGS: -lstdc++
// #include <stdlib.h>
// #include "kextsymboltool.h"
import "C"
import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"unsafe"

	"github.com/blacktop/go-macho"
	"github.com/blacktop/go-macho/pkg/fixupchains"
	"github.com/blacktop/go-macho/types"
	"github.com/blacktop/go-plist"
	"github.com/blacktop/ipsw/pkg/kernelcache"
	"gopkg.in/yaml.v3"

	"github.com/pkg/errors"
)

func stringsToCharPointerPointer(strs []string) **C.char {
	// Allocate an array of character pointers, one for each string in strs.
	chars := make([]*C.char, len(strs))
	for i, s := range strs {
		// Convert each string to a *char and store it in the array.
		chars[i] = C.CString(s)
	}

	// Convert the array of *char pointers to a **char (a pointer to a pointer
	// to a character) and return it.
	return (**C.char)(unsafe.Pointer(&chars[0]))
}

func createSymbolset(imports []string, export string, output string) int {
	args := []string{""}
	for _, str := range imports {
		args = append(args, "-import", str)
	}
	args = append(args, "-export", export, "-output", output)
	charArray := stringsToCharPointerPointer(args)
	defer C.freeChars(C.int(len(args)), charArray)
	return int(C.kextsymboltool(C.int(len(args)), charArray))
}

func allSymbols(path string) ([]string, error) {
	cmd := exec.Command("bash", "-c", fmt.Sprintf("nm -gj \"%s\" | sort -u", path))

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return strings.Split(strings.Trim(out.String(), " \n"), "\n"), nil
}

var symbolStubs []string = []string{
	"__ZN16IOPlatformExpert",
	"__ZN17IONVRAMController",
	"__ZN18IODTPlatformExpert",
	"__ZN24IOCPUInterruptController",
	"__ZN5IOCPU",
	"__ZNK16IOPlatformExpert",
	"__ZNK18IODTPlatformExpert",
	"__ZNK24IOCPUInterruptController",
	"__ZNK5IOCPU",
}

type symbol struct {
	Name   string `plist:"SymbolName,omitempty"`
	Prefix string `plist:"SymbolPrefix,omitempty"`
}

type cFBundle struct {
	ID                string   `plist:"CFBundleIdentifier,omitempty"`
	CompatibleVersion string   `plist:"OSBundleCompatibleVersion,omitempty"`
	Version           string   `plist:"CFBundleVersion,omitempty"`
	Symbols           []symbol `plist:"Symbols,omitempty"`
}

type symbolsSets struct {
	SymbolsSetsDictionary []cFBundle `plist:"SymbolsSets,omitempty"`
}

func extractSymbolsets(path string) (map[string][]string, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("file %s does not exist", path)
	}

	m, err := macho.Open(path)
	if err != nil {
		return nil, errors.Wrapf(err, "%s appears to not be a valid MachO", path)
	}

	if m.FileTOC.FileHeader.Type == types.MH_FILESET {
		m, err = m.GetFileSetFileByName("com.apple.kernel")
		if err != nil {
			return nil, fmt.Errorf("failed to parse entry com.apple.kernel; %v", err)
		}
	}

	symbolsets := m.Section("__LINKINFO", "__symbolsets")

	if symbolsets == nil {
		return nil, fmt.Errorf("kernelcache does NOT contain __LINKINFO.__symbolsets")
	}

	dat := make([]byte, symbolsets.Size)
	m.ReadAt(dat, int64(symbolsets.Offset))

	var blist symbolsSets

	dec := plist.NewDecoder(bytes.NewReader(dat))

	err = dec.Decode(&blist)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse __symbolsets bplist data")
	}

	sets := make(map[string][]string, len(blist.SymbolsSetsDictionary))

	for _, sset := range blist.SymbolsSetsDictionary {
		symbols := make([]string, len(sset.Symbols))
		for i, sym := range sset.Symbols {
			symbols[i] = fmt.Sprintf("%s%s", sym.Prefix, sym.Name)
		}
		sets[sset.ID] = symbols
	}

	return sets, nil
}

type KDKBuilder struct {
	VolumeGroup                string `yaml:"volume_group"`
	KernelPath                 string `yaml:"custom_kernel_path"`
	DestinationKernelcachePath string `yaml:"kernelcache_dst"`
	DestinationKDKPath         string `yaml:"kdk_root"`
	KernelManagement           bool   `yaml:"kernel_management"`
}

func NewKDKBuilderFromPath(path string) (*KDKBuilder, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	builder := KDKBuilder{}
	if err = yaml.Unmarshal(bytes, &builder); err != nil {
		return nil, err
	}
	return &builder, nil
}

type Preboot struct {
	BootConfigurations      map[string][]string
	ActiveBootConfiguration string
}

func NewPrebootFromVolumeGroup(group string) (*Preboot, error) {
	rootPath := fmt.Sprintf("/System/Volumes/Preboot/%s/boot", group)
	contents, err := os.ReadDir(rootPath)
	if err != nil {
		return nil, err
	}
	preboot := Preboot{
		BootConfigurations:      map[string][]string{},
		ActiveBootConfiguration: "",
	}
	for _, entry := range contents {
		switch entry.Name() {
		case "active":
			activeBytes, err := os.ReadFile(path.Join(rootPath,
				entry.Name()))
			if err != nil {
				return nil, err
			}
			preboot.ActiveBootConfiguration = string(activeBytes)
		case "System":
			continue
		default:
			configRootPath := path.Join(rootPath, entry.Name())
			kernelcachePath := path.Join(configRootPath, "System/Library/Caches/com.apple.kernelcaches")
			entries, err := os.ReadDir(kernelcachePath)
			if err != nil {
				return nil, err
			}
			kernelcaches := make([]string, len(entries))
			for i, entry := range entries {
				kernelcaches[i] = path.Join(kernelcachePath, entry.Name())
			}
			preboot.BootConfigurations[entry.Name()] = kernelcaches
		}
	}
	return &preboot, nil
}

func (k *KDKBuilder) LocateOriginalKernelcache() (string, error) {
	preboot, err := NewPrebootFromVolumeGroup(k.VolumeGroup)
	if err != nil {
		return "", err
	}
	if cachePaths, ok := preboot.BootConfigurations[preboot.ActiveBootConfiguration]; ok {
		for _, cachePath := range cachePaths {
			if path.Base(cachePath) == "kernelcache" {
				return cachePath, nil
			}
		}
		if len(cachePaths) > 0 {
			return cachePaths[0], nil
		}
	}
	return "", fmt.Errorf("failed to locate a kernelcache from preboot")
}

func (k *KDKBuilder) KDKKernelCachesPath() string {
	return path.Join(k.DestinationKDKPath, "System/Library/Caches/com.apple.kernelcaches")
}

func (k *KDKBuilder) DecompressedKernelcachePath() string {
	return path.Join(k.KDKKernelCachesPath(), "kernelcache.decompressed")
}

func (k *KDKBuilder) PrepareDecompressedKernelcache() error {
	if err := mkdirp(k.KDKKernelCachesPath()); err != nil {
		return err
	}
	if _, err := os.Stat(k.DecompressedKernelcachePath()); err != nil && !os.IsNotExist(err) {
		return err
	} else if err == nil {
		return nil
	}
	compressed, err := k.LocateOriginalKernelcache()
	if err != nil {
		return err
	}
	if k.KernelManagement {
		return kernelcache.DecompressKernelManagement(compressed, k.DecompressedKernelcachePath())
	} else {
		return kernelcache.Decompress(compressed, k.DecompressedKernelcachePath())
	}
}

func pathExists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, errors.Wrapf(err, "failed to check that %s exists", path)
		}
	}
	return true, nil
}

func mkdirp(path string) error {
	if exists, err := pathExists(path); err != nil {
		return err
	} else if !exists {
		return os.MkdirAll(path, os.FileMode(0777))
	}
	return nil
}

func (k *KDKBuilder) KDKKextRoot() string {
	return path.Join(k.DestinationKDKPath, "System/Library/Extensions")
}

func (k *KDKBuilder) KDKKernelRoot() string {
	return path.Join(k.DestinationKDKPath, "System/Library/Kernels")
}

func (k *KDKBuilder) PrepareKDKStructure() error {
	if err := mkdirp(k.KDKKextRoot()); err != nil {
		return err
	}
	if err := mkdirp(k.KDKKernelRoot()); err != nil {
		return err
	}
	return nil
}

func (k *KDKBuilder) KDKKextPath(bundle kernelcache.CFBundle) string {
	return path.Join(k.DestinationKDKPath, bundle.BundlePath, bundle.RelativePath)
}

func (k *KDKBuilder) PrepareKext(bundle kernelcache.CFBundle) error {
	return mkdirp(path.Dir(k.KDKKextPath(bundle)))
}

func (k *KDKBuilder) ExtractKexts() error {
	if err := k.PrepareKDKStructure(); err != nil {
		return errors.Wrap(err, "failed to prepare kdk structure")
	}
	if err := k.OverlayKDK(); err != nil {
		return err
	}
	if err := k.PrepareDecompressedKernelcache(); err != nil {
		return errors.Wrap(err, "failed to prepare decompressed kernel cache")
	}
	m, err := macho.Open(k.DecompressedKernelcachePath())
	if err != nil {
		return errors.Wrap(err, "failed to open decompressed kernel cache")
	}
	bundles, err := kernelcache.PrelinkInfoDictionaryFromMacho(m)
	if err != nil {
		return errors.Wrap(err, "failed to load bundles from decompressed kernel cache prelink info")
	}
	for _, fe := range m.FileSets() {
		if fe.EntryID == "com.apple.kernel" {
			continue
		}

		bundle, ok := bundles[fe.EntryID]
		if !ok {
			return fmt.Errorf("failed to locate bundle for kext %s", fe.EntryID)
		}
		if err := k.PrepareKext(bundle); err != nil {
			return errors.Wrap(err, "failed to prepare kext")
		}
		kextPath := k.KDKKextPath(bundle)
		if exists, err := pathExists(kextPath); err != nil {
			return errors.Wrapf(err, "failed to check if kext %s exists", fe.EntryID)
		} else if exists {
			continue
		}

		var dcf *fixupchains.DyldChainedFixups
		if m.HasFixups() {
			dcf, err = m.DyldChainedFixups()
			if err != nil {
				return fmt.Errorf("failed to parse fixups from in memory MachO: %v", err)
			}
		}

		baseAddress := m.GetBaseAddress()

		mfe, err := m.GetFileSetFileByName(fe.EntryID)
		if err != nil {
			return fmt.Errorf("failed to parse entry %s: %v", fe.EntryID, err)
		}

		fmt.Printf("extracting %s to %s\n", fe.EntryID, kextPath)
		if err := mfe.Export(kextPath, dcf, baseAddress, nil); err != nil {
			return errors.Wrapf(err, "failed to extract kext %s", fe.EntryID)
		}
	}
	return nil
}

func rsync(src string, dst string) error {
	cmd := exec.Command("bash", "-c", fmt.Sprintf("rsync -rav \"%s\" \"%s\"", src, dst))

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func overlay(src string, dst string) error {
	if err := rsync(src, path.Dir(dst)); err != nil {
		return errors.Wrapf(err, "failed to overlay %s to %s", src, dst)
	}
	return nil
}

func (k *KDKBuilder) OverlayKDK() (err error) {
	err = overlay("/System/Library/Extensions", k.KDKKextRoot())
	if err != nil {
		return
	}
	err = overlay("/System/Library/Kernels", k.KDKKernelRoot())
	if err != nil {
		return
	}
	return
}

const proxyBundlePath = "/System/Library/Extensions/System.kext/PlugIns"

type ProxyBundle struct {
	kernelcache.CFBundle
	KextFolderName string
}

func readProxyBundle(identifier string) (*ProxyBundle, error) {
	entries, err := os.ReadDir(proxyBundlePath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to enumerate System.kext/PlugIns")
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		infoPath := path.Join(proxyBundlePath, entry.Name(), "Info.plist")
		if exists, err := pathExists(infoPath); err != nil {
			return nil, err
		} else if !exists {
			continue
		}
		bytes, err := os.ReadFile(infoPath)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read %s", infoPath)
		}
		bundle := kernelcache.CFBundle{}
		if _, err := plist.Unmarshal(bytes, &bundle); err != nil {
			return nil, err
		}
		if bundle.ID != identifier {
			continue
		}
		return &ProxyBundle{
			bundle,
			entry.Name(),
		}, nil
	}
	return nil, fmt.Errorf("failed to locate Info.plist for %s", identifier)
}

func (k *KDKBuilder) SymbolsetWorkRoot() string {
	return path.Join(k.DestinationKDKPath, "Symbolsets")
}

func (k *KDKBuilder) SymbolsetStubPath() string {
	return path.Join(k.SymbolsetWorkRoot(), "symbolsets_stubs")
}

func (k *KDKBuilder) SymbolsetPath() string {
	return path.Join(k.SymbolsetWorkRoot(), "symbolsets")
}

func (k *KDKBuilder) PrepareSymbolsetRoot() error {
	if err := mkdirp(k.SymbolsetWorkRoot()); err != nil {
		return errors.Wrap(err, "failed to prepare symbolset work root")
	}
	return nil
}

func writeSymbolset(symbols []string, path string) error {
	bytes := []byte(strings.Join(symbols, "\n") + "\n")
	return os.WriteFile(path, bytes, os.FileMode(0666))
}

func (k *KDKBuilder) PrepareSymbolsetStub() error {
	if err := k.PrepareSymbolsetRoot(); err != nil {
		return err
	}
	return writeSymbolset(symbolStubs, k.SymbolsetStubPath())
}

func (k *KDKBuilder) PrepareFullSymbolset() error {
	if err := k.PrepareSymbolsetRoot(); err != nil {
		return err
	}
	symbols, err := allSymbols(k.KernelPath)
	if err != nil {
		return err
	}
	return writeSymbolset(symbols, k.SymbolsetPath())
}

func (k *KDKBuilder) SymbolsetImports() []string {
	return []string{k.SymbolsetStubPath(), k.SymbolsetPath()}
}

func (k *KDKBuilder) PrepareSymbolsetProxies() error {
	if err := k.PrepareSymbolsetStub(); err != nil {
		return err
	}
	if err := k.PrepareFullSymbolset(); err != nil {
		return err
	}
	symbolsets, err := extractSymbolsets(k.KernelPath)
	if err != nil {
		return errors.Wrapf(err, "failed to extract symbolsets from %s", k.KernelPath)
	}
	for id, symbols := range symbolsets {
		bundle, err := readProxyBundle(id)
		if err != nil {
			return err
		}
		kextName := strings.TrimSuffix(bundle.KextFolderName, ".kext")
		symbolsetPath := path.Join(k.SymbolsetWorkRoot(), fmt.Sprintf("%s.exports", kextName))
		if err := writeSymbolset(symbols, symbolsetPath); err != nil {
			return err
		}
		kextPath := path.Join(k.DestinationKDKPath, proxyBundlePath, bundle.KextFolderName)
		executablePath := path.Join(kextPath, bundle.Executable)
		if err := mkdirp(kextPath); err != nil {
			return err
		}
		if code := createSymbolset(k.SymbolsetImports(), symbolsetPath, executablePath); code != 0 {
			return fmt.Errorf("failed to create symbolset binary for %s", kextName)
		}
	}
	return nil
}

func main() {
	builder, err := NewKDKBuilderFromPath("./kdk.local.yaml")
	if err != nil {
		panic(err)
	}
	if err = builder.ExtractKexts(); err != nil {
		panic(err)
	}
	if err = builder.PrepareSymbolsetProxies(); err != nil {
		panic(err)
	}
	return

	fmt.Printf("I'm alive.\n")
	// str := C.CString("")
	if createSymbolset([]string{"allsymbols_stub"}, "com.apple.kpi.private", "Private") != 0 {
		fmt.Printf("failed to create symbolset\n")
	}

	symbols, err := allSymbols("/System/Library/Kernels/kernel")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", len(symbols))

	symbolSets, err := extractSymbolsets("/System/Library/Kernels/Kernel")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", symbolSets)
}
