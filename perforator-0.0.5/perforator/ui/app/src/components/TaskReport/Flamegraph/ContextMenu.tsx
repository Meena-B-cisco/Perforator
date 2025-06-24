import React from 'react';

import { Button, CopyToClipboard, Popup, type PopupProps } from '@gravity-ui/uikit';

import { uiFactory } from 'src/factory';
import type { StringifiedNode } from 'src/models/Profile';


export type PopupData = { offset: [number, number]; node: StringifiedNode };

function getHref(node: StringifiedNode) {
    return uiFactory().goToDefenitionHref(node);
}

type ContextMenuProps = {
    popupData: PopupData;
    anchorRef: PopupProps['anchorRef'];
    onClosePopup: () => void;
};
export const ContextMenu: React.FC<ContextMenuProps> = ({ popupData, anchorRef, onClosePopup }) => {
    const href = getHref(popupData.node);
    const hasFile = Boolean(popupData.node.file);
    const shouldShowGoTo = (
        hasFile &&
        popupData.node.frameOrigin !== 'kernel' &&
        href
    );
    const commonButtonProps = {
        view: 'flat',
        width: 'max',
        onClick: onClosePopup,
        pin: 'brick-brick',
        size: 'l',
    } as const;

    return <Popup
        altBoundary={true}
        open={Boolean(popupData)}
        anchorRef={anchorRef}
        offset={popupData.offset}
        contentClassName={'flamegraph__popup'}
        placement={['top-start']}
        onEscapeKeyDown={onClosePopup}
    >
        <CopyToClipboard text={popupData.node.textId} >
            {() => <Button
                {...commonButtonProps}
            >

            Copy name
            </Button>}
        </CopyToClipboard>
        {/* eslint-disable-next-line no-nested-ternary */}
        {shouldShowGoTo ? (
            <Button
                {...commonButtonProps}
                href={href}
                target="_blank"
            >
                Go to source
            </Button>

        ) : hasFile ? (
            <CopyToClipboard text={popupData.node.file} >
                {() => <Button
                    {...commonButtonProps}
                >
                    Copy file path
                </Button>}
            </CopyToClipboard>
        ) : null}
    </Popup>;
};
