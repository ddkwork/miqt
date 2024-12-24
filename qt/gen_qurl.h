#pragma once
#ifndef MIQT_QT_GEN_QURL_H
#define MIQT_QT_GEN_QURL_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QByteArrayView QByteArrayView;
typedef struct QUrl QUrl;
typedef struct QUrlQuery QUrlQuery;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QUrl* QUrl_new();
extern __declspec(dllexport) QUrl* QUrl_new2(QUrl* copyVal);
extern __declspec(dllexport) QUrl* QUrl_new3(struct miqt_string url);
extern __declspec(dllexport) QUrl* QUrl_new4(struct miqt_string url, ParsingMode mode);
extern __declspec(dllexport) void QUrl_OperatorAssign(QUrl* self, QUrl* copyVal);
extern __declspec(dllexport) void QUrl_OperatorAssignWithUrl(QUrl* self, struct miqt_string url);
extern __declspec(dllexport) void QUrl_Swap(QUrl* self, QUrl* other);
extern __declspec(dllexport) void QUrl_SetUrl(QUrl* self, struct miqt_string url);
extern __declspec(dllexport) struct miqt_string QUrl_Url(const QUrl* self);
extern __declspec(dllexport) struct miqt_string QUrl_ToString(const QUrl* self);
extern __declspec(dllexport) struct miqt_string QUrl_ToDisplayString(const QUrl* self);
extern __declspec(dllexport) QUrl* QUrl_Adjusted(const QUrl* self, FormattingOptions options);
extern __declspec(dllexport) struct miqt_string QUrl_ToEncoded(const QUrl* self);
extern __declspec(dllexport) QUrl* QUrl_FromEncoded(QByteArrayView* input);
extern __declspec(dllexport) QUrl* QUrl_FromUserInput(struct miqt_string userInput);
extern __declspec(dllexport) bool QUrl_IsValid(const QUrl* self);
extern __declspec(dllexport) struct miqt_string QUrl_ErrorString(const QUrl* self);
extern __declspec(dllexport) bool QUrl_IsEmpty(const QUrl* self);
extern __declspec(dllexport) void QUrl_Clear(QUrl* self);
extern __declspec(dllexport) void QUrl_SetScheme(QUrl* self, struct miqt_string scheme);
extern __declspec(dllexport) struct miqt_string QUrl_Scheme(const QUrl* self);
extern __declspec(dllexport) void QUrl_SetAuthority(QUrl* self, struct miqt_string authority);
extern __declspec(dllexport) struct miqt_string QUrl_Authority(const QUrl* self);
extern __declspec(dllexport) void QUrl_SetUserInfo(QUrl* self, struct miqt_string userInfo);
extern __declspec(dllexport) struct miqt_string QUrl_UserInfo(const QUrl* self);
extern __declspec(dllexport) void QUrl_SetUserName(QUrl* self, struct miqt_string userName);
extern __declspec(dllexport) struct miqt_string QUrl_UserName(const QUrl* self);
extern __declspec(dllexport) void QUrl_SetPassword(QUrl* self, struct miqt_string password);
extern __declspec(dllexport) struct miqt_string QUrl_Password(const QUrl* self);
extern __declspec(dllexport) void QUrl_SetHost(QUrl* self, struct miqt_string host);
extern __declspec(dllexport) struct miqt_string QUrl_Host(const QUrl* self);
extern __declspec(dllexport) void QUrl_SetPort(QUrl* self, int port);
extern __declspec(dllexport) int QUrl_Port(const QUrl* self);
extern __declspec(dllexport) void QUrl_SetPath(QUrl* self, struct miqt_string path);
extern __declspec(dllexport) struct miqt_string QUrl_Path(const QUrl* self);
extern __declspec(dllexport) struct miqt_string QUrl_FileName(const QUrl* self);
extern __declspec(dllexport) bool QUrl_HasQuery(const QUrl* self);
extern __declspec(dllexport) void QUrl_SetQuery(QUrl* self, struct miqt_string query);
extern __declspec(dllexport) void QUrl_SetQueryWithQuery(QUrl* self, QUrlQuery* query);
extern __declspec(dllexport) struct miqt_string QUrl_Query(const QUrl* self);
extern __declspec(dllexport) bool QUrl_HasFragment(const QUrl* self);
extern __declspec(dllexport) struct miqt_string QUrl_Fragment(const QUrl* self);
extern __declspec(dllexport) void QUrl_SetFragment(QUrl* self, struct miqt_string fragment);
extern __declspec(dllexport) QUrl* QUrl_Resolved(const QUrl* self, QUrl* relative);
extern __declspec(dllexport) bool QUrl_IsRelative(const QUrl* self);
extern __declspec(dllexport) bool QUrl_IsParentOf(const QUrl* self, QUrl* url);
extern __declspec(dllexport) bool QUrl_IsLocalFile(const QUrl* self);
extern __declspec(dllexport) QUrl* QUrl_FromLocalFile(struct miqt_string localfile);
extern __declspec(dllexport) struct miqt_string QUrl_ToLocalFile(const QUrl* self);
extern __declspec(dllexport) void QUrl_Detach(QUrl* self);
extern __declspec(dllexport) bool QUrl_IsDetached(const QUrl* self);
extern __declspec(dllexport) bool QUrl_Matches(const QUrl* self, QUrl* url, FormattingOptions options);
extern __declspec(dllexport) struct miqt_string QUrl_FromPercentEncoding(struct miqt_string param1);
extern __declspec(dllexport) struct miqt_string QUrl_ToPercentEncoding(struct miqt_string param1);
extern __declspec(dllexport) struct miqt_string QUrl_FromAce(struct miqt_string domain);
extern __declspec(dllexport) struct miqt_string QUrl_ToAce(struct miqt_string domain);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QUrl_IdnWhitelist();
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QUrl_ToStringList(struct miqt_array /* of QUrl* */  uris);
extern __declspec(dllexport) struct miqt_array /* of QUrl* */  QUrl_FromStringList(struct miqt_array /* of struct miqt_string */  uris);
extern __declspec(dllexport) void QUrl_SetIdnWhitelist(struct miqt_array /* of struct miqt_string */  idnWhitelist);
extern __declspec(dllexport) DataPtr* QUrl_DataPtr(QUrl* self);
extern __declspec(dllexport) void QUrl_SetUrl2(QUrl* self, struct miqt_string url, ParsingMode mode);
extern __declspec(dllexport) struct miqt_string QUrl_Url1(const QUrl* self, FormattingOptions options);
extern __declspec(dllexport) struct miqt_string QUrl_ToString1(const QUrl* self, FormattingOptions options);
extern __declspec(dllexport) struct miqt_string QUrl_ToDisplayString1(const QUrl* self, FormattingOptions options);
extern __declspec(dllexport) struct miqt_string QUrl_ToEncoded1(const QUrl* self, FormattingOptions options);
extern __declspec(dllexport) QUrl* QUrl_FromEncoded2(QByteArrayView* input, ParsingMode mode);
extern __declspec(dllexport) QUrl* QUrl_FromUserInput2(struct miqt_string userInput, struct miqt_string workingDirectory);
extern __declspec(dllexport) QUrl* QUrl_FromUserInput3(struct miqt_string userInput, struct miqt_string workingDirectory, UserInputResolutionOptions options);
extern __declspec(dllexport) void QUrl_SetAuthority2(QUrl* self, struct miqt_string authority, ParsingMode mode);
extern __declspec(dllexport) struct miqt_string QUrl_Authority1(const QUrl* self, ComponentFormattingOptions options);
extern __declspec(dllexport) void QUrl_SetUserInfo2(QUrl* self, struct miqt_string userInfo, ParsingMode mode);
extern __declspec(dllexport) struct miqt_string QUrl_UserInfo1(const QUrl* self, ComponentFormattingOptions options);
extern __declspec(dllexport) void QUrl_SetUserName2(QUrl* self, struct miqt_string userName, ParsingMode mode);
extern __declspec(dllexport) struct miqt_string QUrl_UserName1(const QUrl* self, ComponentFormattingOptions options);
extern __declspec(dllexport) void QUrl_SetPassword2(QUrl* self, struct miqt_string password, ParsingMode mode);
extern __declspec(dllexport) struct miqt_string QUrl_Password1(const QUrl* self, ComponentFormattingOptions param1);
extern __declspec(dllexport) void QUrl_SetHost2(QUrl* self, struct miqt_string host, ParsingMode mode);
extern __declspec(dllexport) struct miqt_string QUrl_Host1(const QUrl* self, ComponentFormattingOptions param1);
extern __declspec(dllexport) int QUrl_Port1(const QUrl* self, int defaultPort);
extern __declspec(dllexport) void QUrl_SetPath2(QUrl* self, struct miqt_string path, ParsingMode mode);
extern __declspec(dllexport) struct miqt_string QUrl_Path1(const QUrl* self, ComponentFormattingOptions options);
extern __declspec(dllexport) struct miqt_string QUrl_FileName1(const QUrl* self, ComponentFormattingOptions options);
extern __declspec(dllexport) void QUrl_SetQuery2(QUrl* self, struct miqt_string query, ParsingMode mode);
extern __declspec(dllexport) struct miqt_string QUrl_Query1(const QUrl* self, ComponentFormattingOptions param1);
extern __declspec(dllexport) struct miqt_string QUrl_Fragment1(const QUrl* self, ComponentFormattingOptions options);
extern __declspec(dllexport) void QUrl_SetFragment2(QUrl* self, struct miqt_string fragment, ParsingMode mode);
extern __declspec(dllexport) struct miqt_string QUrl_ToPercentEncoding2(struct miqt_string param1, struct miqt_string exclude);
extern __declspec(dllexport) struct miqt_string QUrl_ToPercentEncoding3(struct miqt_string param1, struct miqt_string exclude, struct miqt_string include);
extern __declspec(dllexport) struct miqt_string QUrl_FromAce2(struct miqt_string domain, AceProcessingOptions options);
extern __declspec(dllexport) struct miqt_string QUrl_ToAce2(struct miqt_string domain, AceProcessingOptions options);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QUrl_ToStringList2(struct miqt_array /* of QUrl* */  uris, FormattingOptions options);
extern __declspec(dllexport) struct miqt_array /* of QUrl* */  QUrl_FromStringList2(struct miqt_array /* of struct miqt_string */  uris, ParsingMode mode);
extern __declspec(dllexport) void QUrl_Delete(QUrl* self, bool isSubclass);

} 
