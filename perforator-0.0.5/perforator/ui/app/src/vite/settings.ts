import type { Alias, BuildOptions, PluginOption } from 'vite';


export interface ViteSettings {
    plugins?: PluginOption[];
    aliases?: Alias[];
    build?: BuildOptions;
}
