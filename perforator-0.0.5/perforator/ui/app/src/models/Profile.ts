type StringTableIndex = number

export interface FormatNode {
    parentIndex: number;
    textId: StringTableIndex;
    sampleCount: number;
    eventCount: number;
    /**
     * either hash once after downloading or get from mapping
     * already darkened if dark theme is active
     */
    color?: string | number;
    /** only add during render */
    x?: number;
    /** only for diff */
    baseEventCount?: number;
    /** only for diff */
    baseSampleCount?: number;
    frameOrigin?: StringTableIndex;
    file?: StringTableIndex;
    kind?: StringTableIndex;
    inlined?: boolean;
}

export type ProfileData = {
    rows: FormatNode[][];
    stringTable: string[];
}

type StringifiableFields = 'frameOrigin' | 'file' | 'kind' | 'textId';
export type StringifiedNode = Omit<FormatNode, StringifiableFields> & {
    [key in StringifiableFields]: string;
};
