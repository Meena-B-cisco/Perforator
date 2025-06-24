from dataclasses import dataclass

from devtools.frontend_build_platform.libraries.logging import timeit

from .base_builder import BaseBuilder
from ..models import CommonBuildersOptions


def touch(file_path: str) -> None:
    with open(file_path, 'w'):
        pass


@dataclass
class PackageBuilderOptions(CommonBuildersOptions):
    pass


class PackageBuilder(BaseBuilder):
    def __init__(self, options):
        super(PackageBuilder, self).__init__(options, copy_package_json=True)

    @timeit
    def bundle(self):
        if self.options.with_after_build and self.options.after_build_outdir:
            return self.bundle_dirs([self.options.after_build_outdir], self.options.bindir, self.options.output_file)
        else:
            touch(self.options.output_file)

    def _build(self):
        pass
