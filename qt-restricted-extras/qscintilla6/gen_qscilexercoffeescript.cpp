#include <QColor>
#include <QFont>
#include <QList>
#include <QMetaObject>
#include <QObject>
#include <QSettings>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qscilexercoffeescript.h>
#include "gen_qscilexercoffeescript.h"

#ifndef _Bool
#define _Bool bool
#endif
#include "_cgo_export.h"

class MiqtVirtualQsciLexerCoffeeScript : public virtual QsciLexerCoffeeScript {
public:

	MiqtVirtualQsciLexerCoffeeScript(): QsciLexerCoffeeScript() {};
	MiqtVirtualQsciLexerCoffeeScript(QObject* parent): QsciLexerCoffeeScript(parent) {};

	virtual ~MiqtVirtualQsciLexerCoffeeScript() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Language = 0;

	// Subclass to allow providing a Go implementation
	virtual const char* language() const override {
		if (handle__Language == 0) {
			return nullptr; // Pure virtual, there is no base we can call
		}
		

		const char* callback_return_value = miqt_exec_callback_QsciLexerCoffeeScript_Language(const_cast<MiqtVirtualQsciLexerCoffeeScript*>(this), handle__Language);

		return callback_return_value;
	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Lexer = 0;

	// Subclass to allow providing a Go implementation
	virtual const char* lexer() const override {
		if (handle__Lexer == 0) {
			return QsciLexerCoffeeScript::lexer();
		}
		

		const char* callback_return_value = miqt_exec_callback_QsciLexerCoffeeScript_Lexer(const_cast<MiqtVirtualQsciLexerCoffeeScript*>(this), handle__Lexer);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	const char* virtualbase_Lexer() const {

		return (const char*) QsciLexerCoffeeScript::lexer();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__LexerId = 0;

	// Subclass to allow providing a Go implementation
	virtual int lexerId() const override {
		if (handle__LexerId == 0) {
			return QsciLexerCoffeeScript::lexerId();
		}
		

		int callback_return_value = miqt_exec_callback_QsciLexerCoffeeScript_LexerId(const_cast<MiqtVirtualQsciLexerCoffeeScript*>(this), handle__LexerId);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_LexerId() const {

		return QsciLexerCoffeeScript::lexerId();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__AutoCompletionFillups = 0;

	// Subclass to allow providing a Go implementation
	virtual const char* autoCompletionFillups() const override {
		if (handle__AutoCompletionFillups == 0) {
			return QsciLexerCoffeeScript::autoCompletionFillups();
		}
		

		const char* callback_return_value = miqt_exec_callback_QsciLexerCoffeeScript_AutoCompletionFillups(const_cast<MiqtVirtualQsciLexerCoffeeScript*>(this), handle__AutoCompletionFillups);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	const char* virtualbase_AutoCompletionFillups() const {

		return (const char*) QsciLexerCoffeeScript::autoCompletionFillups();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__AutoCompletionWordSeparators = 0;

	// Subclass to allow providing a Go implementation
	virtual QStringList autoCompletionWordSeparators() const override {
		if (handle__AutoCompletionWordSeparators == 0) {
			return QsciLexerCoffeeScript::autoCompletionWordSeparators();
		}
		

		struct miqt_array /* of struct miqt_string */  callback_return_value = miqt_exec_callback_QsciLexerCoffeeScript_AutoCompletionWordSeparators(const_cast<MiqtVirtualQsciLexerCoffeeScript*>(this), handle__AutoCompletionWordSeparators);
		QStringList callback_return_value_QList;
		callback_return_value_QList.reserve(callback_return_value.len);
		struct miqt_string* callback_return_value_arr = static_cast<struct miqt_string*>(callback_return_value.data);
		for(size_t i = 0; i < callback_return_value.len; ++i) {
			QString callback_return_value_arr_i_QString = QString::fromUtf8(callback_return_value_arr[i].data, callback_return_value_arr[i].len);
			callback_return_value_QList.push_back(callback_return_value_arr_i_QString);
		}

		return callback_return_value_QList;
	}

	// Wrapper to allow calling protected method
	struct miqt_array /* of struct miqt_string */  virtualbase_AutoCompletionWordSeparators() const {

		QStringList _ret = QsciLexerCoffeeScript::autoCompletionWordSeparators();
		// Convert QList<> from C++ memory to manually-managed C memory
		struct miqt_string* _arr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.length()));
		for (size_t i = 0, e = _ret.length(); i < e; ++i) {
			QString _lv_ret = _ret[i];
			// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
			QByteArray _lv_b = _lv_ret.toUtf8();
			struct miqt_string _lv_ms;
			_lv_ms.len = _lv_b.length();
			_lv_ms.data = static_cast<char*>(malloc(_lv_ms.len));
			memcpy(_lv_ms.data, _lv_b.data(), _lv_ms.len);
			_arr[i] = _lv_ms;
		}
		struct miqt_array _out;
		_out.len = _ret.length();
		_out.data = static_cast<void*>(_arr);
		return _out;

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__BlockEnd = 0;

	// Subclass to allow providing a Go implementation
	virtual const char* blockEnd(int* style) const override {
		if (handle__BlockEnd == 0) {
			return QsciLexerCoffeeScript::blockEnd(style);
		}
		
		int* sigval1 = style;

		const char* callback_return_value = miqt_exec_callback_QsciLexerCoffeeScript_BlockEnd(const_cast<MiqtVirtualQsciLexerCoffeeScript*>(this), handle__BlockEnd, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	const char* virtualbase_BlockEnd(int* style) const {

		return (const char*) QsciLexerCoffeeScript::blockEnd(static_cast<int*>(style));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__BlockLookback = 0;

	// Subclass to allow providing a Go implementation
	virtual int blockLookback() const override {
		if (handle__BlockLookback == 0) {
			return QsciLexerCoffeeScript::blockLookback();
		}
		

		int callback_return_value = miqt_exec_callback_QsciLexerCoffeeScript_BlockLookback(const_cast<MiqtVirtualQsciLexerCoffeeScript*>(this), handle__BlockLookback);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_BlockLookback() const {

		return QsciLexerCoffeeScript::blockLookback();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__BlockStart = 0;

	// Subclass to allow providing a Go implementation
	virtual const char* blockStart(int* style) const override {
		if (handle__BlockStart == 0) {
			return QsciLexerCoffeeScript::blockStart(style);
		}
		
		int* sigval1 = style;

		const char* callback_return_value = miqt_exec_callback_QsciLexerCoffeeScript_BlockStart(const_cast<MiqtVirtualQsciLexerCoffeeScript*>(this), handle__BlockStart, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	const char* virtualbase_BlockStart(int* style) const {

		return (const char*) QsciLexerCoffeeScript::blockStart(static_cast<int*>(style));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__BlockStartKeyword = 0;

	// Subclass to allow providing a Go implementation
	virtual const char* blockStartKeyword(int* style) const override {
		if (handle__BlockStartKeyword == 0) {
			return QsciLexerCoffeeScript::blockStartKeyword(style);
		}
		
		int* sigval1 = style;

		const char* callback_return_value = miqt_exec_callback_QsciLexerCoffeeScript_BlockStartKeyword(const_cast<MiqtVirtualQsciLexerCoffeeScript*>(this), handle__BlockStartKeyword, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	const char* virtualbase_BlockStartKeyword(int* style) const {

		return (const char*) QsciLexerCoffeeScript::blockStartKeyword(static_cast<int*>(style));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__BraceStyle = 0;

	// Subclass to allow providing a Go implementation
	virtual int braceStyle() const override {
		if (handle__BraceStyle == 0) {
			return QsciLexerCoffeeScript::braceStyle();
		}
		

		int callback_return_value = miqt_exec_callback_QsciLexerCoffeeScript_BraceStyle(const_cast<MiqtVirtualQsciLexerCoffeeScript*>(this), handle__BraceStyle);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_BraceStyle() const {

		return QsciLexerCoffeeScript::braceStyle();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__CaseSensitive = 0;

	// Subclass to allow providing a Go implementation
	virtual bool caseSensitive() const override {
		if (handle__CaseSensitive == 0) {
			return QsciLexerCoffeeScript::caseSensitive();
		}
		

		bool callback_return_value = miqt_exec_callback_QsciLexerCoffeeScript_CaseSensitive(const_cast<MiqtVirtualQsciLexerCoffeeScript*>(this), handle__CaseSensitive);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_CaseSensitive() const {

		return QsciLexerCoffeeScript::caseSensitive();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Color = 0;

	// Subclass to allow providing a Go implementation
	virtual QColor color(int style) const override {
		if (handle__Color == 0) {
			return QsciLexerCoffeeScript::color(style);
		}
		
		int sigval1 = style;

		QColor* callback_return_value = miqt_exec_callback_QsciLexerCoffeeScript_Color(const_cast<MiqtVirtualQsciLexerCoffeeScript*>(this), handle__Color, sigval1);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QColor* virtualbase_Color(int style) const {

		return new QColor(QsciLexerCoffeeScript::color(static_cast<int>(style)));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__EolFill = 0;

	// Subclass to allow providing a Go implementation
	virtual bool eolFill(int style) const override {
		if (handle__EolFill == 0) {
			return QsciLexerCoffeeScript::eolFill(style);
		}
		
		int sigval1 = style;

		bool callback_return_value = miqt_exec_callback_QsciLexerCoffeeScript_EolFill(const_cast<MiqtVirtualQsciLexerCoffeeScript*>(this), handle__EolFill, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_EolFill(int style) const {

		return QsciLexerCoffeeScript::eolFill(static_cast<int>(style));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Font = 0;

	// Subclass to allow providing a Go implementation
	virtual QFont font(int style) const override {
		if (handle__Font == 0) {
			return QsciLexerCoffeeScript::font(style);
		}
		
		int sigval1 = style;

		QFont* callback_return_value = miqt_exec_callback_QsciLexerCoffeeScript_Font(const_cast<MiqtVirtualQsciLexerCoffeeScript*>(this), handle__Font, sigval1);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QFont* virtualbase_Font(int style) const {

		return new QFont(QsciLexerCoffeeScript::font(static_cast<int>(style)));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__IndentationGuideView = 0;

	// Subclass to allow providing a Go implementation
	virtual int indentationGuideView() const override {
		if (handle__IndentationGuideView == 0) {
			return QsciLexerCoffeeScript::indentationGuideView();
		}
		

		int callback_return_value = miqt_exec_callback_QsciLexerCoffeeScript_IndentationGuideView(const_cast<MiqtVirtualQsciLexerCoffeeScript*>(this), handle__IndentationGuideView);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_IndentationGuideView() const {

		return QsciLexerCoffeeScript::indentationGuideView();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Keywords = 0;

	// Subclass to allow providing a Go implementation
	virtual const char* keywords(int set) const override {
		if (handle__Keywords == 0) {
			return QsciLexerCoffeeScript::keywords(set);
		}
		
		int sigval1 = set;

		const char* callback_return_value = miqt_exec_callback_QsciLexerCoffeeScript_Keywords(const_cast<MiqtVirtualQsciLexerCoffeeScript*>(this), handle__Keywords, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	const char* virtualbase_Keywords(int set) const {

		return (const char*) QsciLexerCoffeeScript::keywords(static_cast<int>(set));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DefaultStyle = 0;

	// Subclass to allow providing a Go implementation
	virtual int defaultStyle() const override {
		if (handle__DefaultStyle == 0) {
			return QsciLexerCoffeeScript::defaultStyle();
		}
		

		int callback_return_value = miqt_exec_callback_QsciLexerCoffeeScript_DefaultStyle(const_cast<MiqtVirtualQsciLexerCoffeeScript*>(this), handle__DefaultStyle);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_DefaultStyle() const {

		return QsciLexerCoffeeScript::defaultStyle();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Description = 0;

	// Subclass to allow providing a Go implementation
	virtual QString description(int style) const override {
		if (handle__Description == 0) {
			return QString(); // Pure virtual, there is no base we can call
		}
		
		int sigval1 = style;

		struct miqt_string callback_return_value = miqt_exec_callback_QsciLexerCoffeeScript_Description(const_cast<MiqtVirtualQsciLexerCoffeeScript*>(this), handle__Description, sigval1);
		QString callback_return_value_QString = QString::fromUtf8(callback_return_value.data, callback_return_value.len);

		return callback_return_value_QString;
	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Paper = 0;

	// Subclass to allow providing a Go implementation
	virtual QColor paper(int style) const override {
		if (handle__Paper == 0) {
			return QsciLexerCoffeeScript::paper(style);
		}
		
		int sigval1 = style;

		QColor* callback_return_value = miqt_exec_callback_QsciLexerCoffeeScript_Paper(const_cast<MiqtVirtualQsciLexerCoffeeScript*>(this), handle__Paper, sigval1);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QColor* virtualbase_Paper(int style) const {

		return new QColor(QsciLexerCoffeeScript::paper(static_cast<int>(style)));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DefaultColorWithStyle = 0;

	// Subclass to allow providing a Go implementation
	virtual QColor defaultColor(int style) const override {
		if (handle__DefaultColorWithStyle == 0) {
			return QsciLexerCoffeeScript::defaultColor(style);
		}
		
		int sigval1 = style;

		QColor* callback_return_value = miqt_exec_callback_QsciLexerCoffeeScript_DefaultColorWithStyle(const_cast<MiqtVirtualQsciLexerCoffeeScript*>(this), handle__DefaultColorWithStyle, sigval1);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QColor* virtualbase_DefaultColorWithStyle(int style) const {

		return new QColor(QsciLexerCoffeeScript::defaultColor(static_cast<int>(style)));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DefaultEolFill = 0;

	// Subclass to allow providing a Go implementation
	virtual bool defaultEolFill(int style) const override {
		if (handle__DefaultEolFill == 0) {
			return QsciLexerCoffeeScript::defaultEolFill(style);
		}
		
		int sigval1 = style;

		bool callback_return_value = miqt_exec_callback_QsciLexerCoffeeScript_DefaultEolFill(const_cast<MiqtVirtualQsciLexerCoffeeScript*>(this), handle__DefaultEolFill, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_DefaultEolFill(int style) const {

		return QsciLexerCoffeeScript::defaultEolFill(static_cast<int>(style));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DefaultFontWithStyle = 0;

	// Subclass to allow providing a Go implementation
	virtual QFont defaultFont(int style) const override {
		if (handle__DefaultFontWithStyle == 0) {
			return QsciLexerCoffeeScript::defaultFont(style);
		}
		
		int sigval1 = style;

		QFont* callback_return_value = miqt_exec_callback_QsciLexerCoffeeScript_DefaultFontWithStyle(const_cast<MiqtVirtualQsciLexerCoffeeScript*>(this), handle__DefaultFontWithStyle, sigval1);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QFont* virtualbase_DefaultFontWithStyle(int style) const {

		return new QFont(QsciLexerCoffeeScript::defaultFont(static_cast<int>(style)));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DefaultPaperWithStyle = 0;

	// Subclass to allow providing a Go implementation
	virtual QColor defaultPaper(int style) const override {
		if (handle__DefaultPaperWithStyle == 0) {
			return QsciLexerCoffeeScript::defaultPaper(style);
		}
		
		int sigval1 = style;

		QColor* callback_return_value = miqt_exec_callback_QsciLexerCoffeeScript_DefaultPaperWithStyle(const_cast<MiqtVirtualQsciLexerCoffeeScript*>(this), handle__DefaultPaperWithStyle, sigval1);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QColor* virtualbase_DefaultPaperWithStyle(int style) const {

		return new QColor(QsciLexerCoffeeScript::defaultPaper(static_cast<int>(style)));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetEditor = 0;

	// Subclass to allow providing a Go implementation
	virtual void setEditor(QsciScintilla* editor) override {
		if (handle__SetEditor == 0) {
			QsciLexerCoffeeScript::setEditor(editor);
			return;
		}
		
		QsciScintilla* sigval1 = editor;

		miqt_exec_callback_QsciLexerCoffeeScript_SetEditor(this, handle__SetEditor, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SetEditor(QsciScintilla* editor) {

		QsciLexerCoffeeScript::setEditor(editor);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__RefreshProperties = 0;

	// Subclass to allow providing a Go implementation
	virtual void refreshProperties() override {
		if (handle__RefreshProperties == 0) {
			QsciLexerCoffeeScript::refreshProperties();
			return;
		}
		

		miqt_exec_callback_QsciLexerCoffeeScript_RefreshProperties(this, handle__RefreshProperties);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_RefreshProperties() {

		QsciLexerCoffeeScript::refreshProperties();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__StyleBitsNeeded = 0;

	// Subclass to allow providing a Go implementation
	virtual int styleBitsNeeded() const override {
		if (handle__StyleBitsNeeded == 0) {
			return QsciLexerCoffeeScript::styleBitsNeeded();
		}
		

		int callback_return_value = miqt_exec_callback_QsciLexerCoffeeScript_StyleBitsNeeded(const_cast<MiqtVirtualQsciLexerCoffeeScript*>(this), handle__StyleBitsNeeded);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_StyleBitsNeeded() const {

		return QsciLexerCoffeeScript::styleBitsNeeded();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__WordCharacters = 0;

	// Subclass to allow providing a Go implementation
	virtual const char* wordCharacters() const override {
		if (handle__WordCharacters == 0) {
			return QsciLexerCoffeeScript::wordCharacters();
		}
		

		const char* callback_return_value = miqt_exec_callback_QsciLexerCoffeeScript_WordCharacters(const_cast<MiqtVirtualQsciLexerCoffeeScript*>(this), handle__WordCharacters);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	const char* virtualbase_WordCharacters() const {

		return (const char*) QsciLexerCoffeeScript::wordCharacters();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetAutoIndentStyle = 0;

	// Subclass to allow providing a Go implementation
	virtual void setAutoIndentStyle(int autoindentstyle) override {
		if (handle__SetAutoIndentStyle == 0) {
			QsciLexerCoffeeScript::setAutoIndentStyle(autoindentstyle);
			return;
		}
		
		int sigval1 = autoindentstyle;

		miqt_exec_callback_QsciLexerCoffeeScript_SetAutoIndentStyle(this, handle__SetAutoIndentStyle, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SetAutoIndentStyle(int autoindentstyle) {

		QsciLexerCoffeeScript::setAutoIndentStyle(static_cast<int>(autoindentstyle));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetColor = 0;

	// Subclass to allow providing a Go implementation
	virtual void setColor(const QColor& c, int style) override {
		if (handle__SetColor == 0) {
			QsciLexerCoffeeScript::setColor(c, style);
			return;
		}
		
		const QColor& c_ret = c;
		// Cast returned reference into pointer
		QColor* sigval1 = const_cast<QColor*>(&c_ret);
		int sigval2 = style;

		miqt_exec_callback_QsciLexerCoffeeScript_SetColor(this, handle__SetColor, sigval1, sigval2);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SetColor(QColor* c, int style) {

		QsciLexerCoffeeScript::setColor(*c, static_cast<int>(style));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetEolFill = 0;

	// Subclass to allow providing a Go implementation
	virtual void setEolFill(bool eoffill, int style) override {
		if (handle__SetEolFill == 0) {
			QsciLexerCoffeeScript::setEolFill(eoffill, style);
			return;
		}
		
		bool sigval1 = eoffill;
		int sigval2 = style;

		miqt_exec_callback_QsciLexerCoffeeScript_SetEolFill(this, handle__SetEolFill, sigval1, sigval2);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SetEolFill(bool eoffill, int style) {

		QsciLexerCoffeeScript::setEolFill(eoffill, static_cast<int>(style));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetFont = 0;

	// Subclass to allow providing a Go implementation
	virtual void setFont(const QFont& f, int style) override {
		if (handle__SetFont == 0) {
			QsciLexerCoffeeScript::setFont(f, style);
			return;
		}
		
		const QFont& f_ret = f;
		// Cast returned reference into pointer
		QFont* sigval1 = const_cast<QFont*>(&f_ret);
		int sigval2 = style;

		miqt_exec_callback_QsciLexerCoffeeScript_SetFont(this, handle__SetFont, sigval1, sigval2);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SetFont(QFont* f, int style) {

		QsciLexerCoffeeScript::setFont(*f, static_cast<int>(style));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SetPaper = 0;

	// Subclass to allow providing a Go implementation
	virtual void setPaper(const QColor& c, int style) override {
		if (handle__SetPaper == 0) {
			QsciLexerCoffeeScript::setPaper(c, style);
			return;
		}
		
		const QColor& c_ret = c;
		// Cast returned reference into pointer
		QColor* sigval1 = const_cast<QColor*>(&c_ret);
		int sigval2 = style;

		miqt_exec_callback_QsciLexerCoffeeScript_SetPaper(this, handle__SetPaper, sigval1, sigval2);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_SetPaper(QColor* c, int style) {

		QsciLexerCoffeeScript::setPaper(*c, static_cast<int>(style));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ReadProperties = 0;

	// Subclass to allow providing a Go implementation
	virtual bool readProperties(QSettings& qs, const QString& prefix) override {
		if (handle__ReadProperties == 0) {
			return QsciLexerCoffeeScript::readProperties(qs, prefix);
		}
		
		QSettings& qs_ret = qs;
		// Cast returned reference into pointer
		QSettings* sigval1 = &qs_ret;
		const QString prefix_ret = prefix;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray prefix_b = prefix_ret.toUtf8();
		struct miqt_string prefix_ms;
		prefix_ms.len = prefix_b.length();
		prefix_ms.data = static_cast<char*>(malloc(prefix_ms.len));
		memcpy(prefix_ms.data, prefix_b.data(), prefix_ms.len);
		struct miqt_string sigval2 = prefix_ms;

		bool callback_return_value = miqt_exec_callback_QsciLexerCoffeeScript_ReadProperties(this, handle__ReadProperties, sigval1, sigval2);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_ReadProperties(QSettings* qs, struct miqt_string prefix) {
		QString prefix_QString = QString::fromUtf8(prefix.data, prefix.len);

		return QsciLexerCoffeeScript::readProperties(*qs, prefix_QString);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__WriteProperties = 0;

	// Subclass to allow providing a Go implementation
	virtual bool writeProperties(QSettings& qs, const QString& prefix) const override {
		if (handle__WriteProperties == 0) {
			return QsciLexerCoffeeScript::writeProperties(qs, prefix);
		}
		
		QSettings& qs_ret = qs;
		// Cast returned reference into pointer
		QSettings* sigval1 = &qs_ret;
		const QString prefix_ret = prefix;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray prefix_b = prefix_ret.toUtf8();
		struct miqt_string prefix_ms;
		prefix_ms.len = prefix_b.length();
		prefix_ms.data = static_cast<char*>(malloc(prefix_ms.len));
		memcpy(prefix_ms.data, prefix_b.data(), prefix_ms.len);
		struct miqt_string sigval2 = prefix_ms;

		bool callback_return_value = miqt_exec_callback_QsciLexerCoffeeScript_WriteProperties(const_cast<MiqtVirtualQsciLexerCoffeeScript*>(this), handle__WriteProperties, sigval1, sigval2);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	bool virtualbase_WriteProperties(QSettings* qs, struct miqt_string prefix) const {
		QString prefix_QString = QString::fromUtf8(prefix.data, prefix.len);

		return QsciLexerCoffeeScript::writeProperties(*qs, prefix_QString);

	}

};

QsciLexerCoffeeScript* QsciLexerCoffeeScript_new() {
	return new MiqtVirtualQsciLexerCoffeeScript();
}

QsciLexerCoffeeScript* QsciLexerCoffeeScript_new2(QObject* parent) {
	return new MiqtVirtualQsciLexerCoffeeScript(parent);
}

void QsciLexerCoffeeScript_virtbase(QsciLexerCoffeeScript* src, QsciLexer** outptr_QsciLexer) {
	*outptr_QsciLexer = static_cast<QsciLexer*>(src);
}

QMetaObject* QsciLexerCoffeeScript_MetaObject(const QsciLexerCoffeeScript* self) {
	return (QMetaObject*) self->metaObject();
}

void* QsciLexerCoffeeScript_Metacast(QsciLexerCoffeeScript* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QsciLexerCoffeeScript_Tr(const char* s) {
	QString _ret = QsciLexerCoffeeScript::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

const char* QsciLexerCoffeeScript_Language(const QsciLexerCoffeeScript* self) {
	return (const char*) self->language();
}

const char* QsciLexerCoffeeScript_Lexer(const QsciLexerCoffeeScript* self) {
	return (const char*) self->lexer();
}

struct miqt_array /* of struct miqt_string */  QsciLexerCoffeeScript_AutoCompletionWordSeparators(const QsciLexerCoffeeScript* self) {
	QStringList _ret = self->autoCompletionWordSeparators();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string* _arr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		struct miqt_string _lv_ms;
		_lv_ms.len = _lv_b.length();
		_lv_ms.data = static_cast<char*>(malloc(_lv_ms.len));
		memcpy(_lv_ms.data, _lv_b.data(), _lv_ms.len);
		_arr[i] = _lv_ms;
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

const char* QsciLexerCoffeeScript_BlockEnd(const QsciLexerCoffeeScript* self) {
	return (const char*) self->blockEnd();
}

const char* QsciLexerCoffeeScript_BlockStart(const QsciLexerCoffeeScript* self) {
	return (const char*) self->blockStart();
}

const char* QsciLexerCoffeeScript_BlockStartKeyword(const QsciLexerCoffeeScript* self) {
	return (const char*) self->blockStartKeyword();
}

int QsciLexerCoffeeScript_BraceStyle(const QsciLexerCoffeeScript* self) {
	return self->braceStyle();
}

const char* QsciLexerCoffeeScript_WordCharacters(const QsciLexerCoffeeScript* self) {
	return (const char*) self->wordCharacters();
}

QColor* QsciLexerCoffeeScript_DefaultColor(const QsciLexerCoffeeScript* self, int style) {
	return new QColor(self->defaultColor(static_cast<int>(style)));
}

bool QsciLexerCoffeeScript_DefaultEolFill(const QsciLexerCoffeeScript* self, int style) {
	return self->defaultEolFill(static_cast<int>(style));
}

QFont* QsciLexerCoffeeScript_DefaultFont(const QsciLexerCoffeeScript* self, int style) {
	return new QFont(self->defaultFont(static_cast<int>(style)));
}

QColor* QsciLexerCoffeeScript_DefaultPaper(const QsciLexerCoffeeScript* self, int style) {
	return new QColor(self->defaultPaper(static_cast<int>(style)));
}

const char* QsciLexerCoffeeScript_Keywords(const QsciLexerCoffeeScript* self, int set) {
	return (const char*) self->keywords(static_cast<int>(set));
}

struct miqt_string QsciLexerCoffeeScript_Description(const QsciLexerCoffeeScript* self, int style) {
	QString _ret = self->description(static_cast<int>(style));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QsciLexerCoffeeScript_RefreshProperties(QsciLexerCoffeeScript* self) {
	self->refreshProperties();
}

bool QsciLexerCoffeeScript_DollarsAllowed(const QsciLexerCoffeeScript* self) {
	return self->dollarsAllowed();
}

void QsciLexerCoffeeScript_SetDollarsAllowed(QsciLexerCoffeeScript* self, bool allowed) {
	self->setDollarsAllowed(allowed);
}

bool QsciLexerCoffeeScript_FoldComments(const QsciLexerCoffeeScript* self) {
	return self->foldComments();
}

void QsciLexerCoffeeScript_SetFoldComments(QsciLexerCoffeeScript* self, bool fold) {
	self->setFoldComments(fold);
}

bool QsciLexerCoffeeScript_FoldCompact(const QsciLexerCoffeeScript* self) {
	return self->foldCompact();
}

void QsciLexerCoffeeScript_SetFoldCompact(QsciLexerCoffeeScript* self, bool fold) {
	self->setFoldCompact(fold);
}

bool QsciLexerCoffeeScript_StylePreprocessor(const QsciLexerCoffeeScript* self) {
	return self->stylePreprocessor();
}

void QsciLexerCoffeeScript_SetStylePreprocessor(QsciLexerCoffeeScript* self, bool style) {
	self->setStylePreprocessor(style);
}

struct miqt_string QsciLexerCoffeeScript_Tr2(const char* s, const char* c) {
	QString _ret = QsciLexerCoffeeScript::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QsciLexerCoffeeScript_Tr3(const char* s, const char* c, int n) {
	QString _ret = QsciLexerCoffeeScript::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

const char* QsciLexerCoffeeScript_BlockEnd1(const QsciLexerCoffeeScript* self, int* style) {
	return (const char*) self->blockEnd(static_cast<int*>(style));
}

const char* QsciLexerCoffeeScript_BlockStart1(const QsciLexerCoffeeScript* self, int* style) {
	return (const char*) self->blockStart(static_cast<int*>(style));
}

const char* QsciLexerCoffeeScript_BlockStartKeyword1(const QsciLexerCoffeeScript* self, int* style) {
	return (const char*) self->blockStartKeyword(static_cast<int*>(style));
}

void QsciLexerCoffeeScript_override_virtual_Language(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__Language = slot;
}

void QsciLexerCoffeeScript_override_virtual_Lexer(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__Lexer = slot;
}

const char* QsciLexerCoffeeScript_virtualbase_Lexer(const void* self) {
	return ( (const MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_Lexer();
}

void QsciLexerCoffeeScript_override_virtual_LexerId(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__LexerId = slot;
}

int QsciLexerCoffeeScript_virtualbase_LexerId(const void* self) {
	return ( (const MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_LexerId();
}

void QsciLexerCoffeeScript_override_virtual_AutoCompletionFillups(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__AutoCompletionFillups = slot;
}

const char* QsciLexerCoffeeScript_virtualbase_AutoCompletionFillups(const void* self) {
	return ( (const MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_AutoCompletionFillups();
}

void QsciLexerCoffeeScript_override_virtual_AutoCompletionWordSeparators(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__AutoCompletionWordSeparators = slot;
}

struct miqt_array /* of struct miqt_string */  QsciLexerCoffeeScript_virtualbase_AutoCompletionWordSeparators(const void* self) {
	return ( (const MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_AutoCompletionWordSeparators();
}

void QsciLexerCoffeeScript_override_virtual_BlockEnd(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__BlockEnd = slot;
}

const char* QsciLexerCoffeeScript_virtualbase_BlockEnd(const void* self, int* style) {
	return ( (const MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_BlockEnd(style);
}

void QsciLexerCoffeeScript_override_virtual_BlockLookback(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__BlockLookback = slot;
}

int QsciLexerCoffeeScript_virtualbase_BlockLookback(const void* self) {
	return ( (const MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_BlockLookback();
}

void QsciLexerCoffeeScript_override_virtual_BlockStart(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__BlockStart = slot;
}

const char* QsciLexerCoffeeScript_virtualbase_BlockStart(const void* self, int* style) {
	return ( (const MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_BlockStart(style);
}

void QsciLexerCoffeeScript_override_virtual_BlockStartKeyword(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__BlockStartKeyword = slot;
}

const char* QsciLexerCoffeeScript_virtualbase_BlockStartKeyword(const void* self, int* style) {
	return ( (const MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_BlockStartKeyword(style);
}

void QsciLexerCoffeeScript_override_virtual_BraceStyle(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__BraceStyle = slot;
}

int QsciLexerCoffeeScript_virtualbase_BraceStyle(const void* self) {
	return ( (const MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_BraceStyle();
}

void QsciLexerCoffeeScript_override_virtual_CaseSensitive(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__CaseSensitive = slot;
}

bool QsciLexerCoffeeScript_virtualbase_CaseSensitive(const void* self) {
	return ( (const MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_CaseSensitive();
}

void QsciLexerCoffeeScript_override_virtual_Color(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__Color = slot;
}

QColor* QsciLexerCoffeeScript_virtualbase_Color(const void* self, int style) {
	return ( (const MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_Color(style);
}

void QsciLexerCoffeeScript_override_virtual_EolFill(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__EolFill = slot;
}

bool QsciLexerCoffeeScript_virtualbase_EolFill(const void* self, int style) {
	return ( (const MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_EolFill(style);
}

void QsciLexerCoffeeScript_override_virtual_Font(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__Font = slot;
}

QFont* QsciLexerCoffeeScript_virtualbase_Font(const void* self, int style) {
	return ( (const MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_Font(style);
}

void QsciLexerCoffeeScript_override_virtual_IndentationGuideView(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__IndentationGuideView = slot;
}

int QsciLexerCoffeeScript_virtualbase_IndentationGuideView(const void* self) {
	return ( (const MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_IndentationGuideView();
}

void QsciLexerCoffeeScript_override_virtual_Keywords(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__Keywords = slot;
}

const char* QsciLexerCoffeeScript_virtualbase_Keywords(const void* self, int set) {
	return ( (const MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_Keywords(set);
}

void QsciLexerCoffeeScript_override_virtual_DefaultStyle(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__DefaultStyle = slot;
}

int QsciLexerCoffeeScript_virtualbase_DefaultStyle(const void* self) {
	return ( (const MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_DefaultStyle();
}

void QsciLexerCoffeeScript_override_virtual_Description(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__Description = slot;
}

void QsciLexerCoffeeScript_override_virtual_Paper(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__Paper = slot;
}

QColor* QsciLexerCoffeeScript_virtualbase_Paper(const void* self, int style) {
	return ( (const MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_Paper(style);
}

void QsciLexerCoffeeScript_override_virtual_DefaultColorWithStyle(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__DefaultColorWithStyle = slot;
}

QColor* QsciLexerCoffeeScript_virtualbase_DefaultColorWithStyle(const void* self, int style) {
	return ( (const MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_DefaultColorWithStyle(style);
}

void QsciLexerCoffeeScript_override_virtual_DefaultEolFill(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__DefaultEolFill = slot;
}

bool QsciLexerCoffeeScript_virtualbase_DefaultEolFill(const void* self, int style) {
	return ( (const MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_DefaultEolFill(style);
}

void QsciLexerCoffeeScript_override_virtual_DefaultFontWithStyle(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__DefaultFontWithStyle = slot;
}

QFont* QsciLexerCoffeeScript_virtualbase_DefaultFontWithStyle(const void* self, int style) {
	return ( (const MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_DefaultFontWithStyle(style);
}

void QsciLexerCoffeeScript_override_virtual_DefaultPaperWithStyle(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__DefaultPaperWithStyle = slot;
}

QColor* QsciLexerCoffeeScript_virtualbase_DefaultPaperWithStyle(const void* self, int style) {
	return ( (const MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_DefaultPaperWithStyle(style);
}

void QsciLexerCoffeeScript_override_virtual_SetEditor(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__SetEditor = slot;
}

void QsciLexerCoffeeScript_virtualbase_SetEditor(void* self, QsciScintilla* editor) {
	( (MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_SetEditor(editor);
}

void QsciLexerCoffeeScript_override_virtual_RefreshProperties(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__RefreshProperties = slot;
}

void QsciLexerCoffeeScript_virtualbase_RefreshProperties(void* self) {
	( (MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_RefreshProperties();
}

void QsciLexerCoffeeScript_override_virtual_StyleBitsNeeded(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__StyleBitsNeeded = slot;
}

int QsciLexerCoffeeScript_virtualbase_StyleBitsNeeded(const void* self) {
	return ( (const MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_StyleBitsNeeded();
}

void QsciLexerCoffeeScript_override_virtual_WordCharacters(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__WordCharacters = slot;
}

const char* QsciLexerCoffeeScript_virtualbase_WordCharacters(const void* self) {
	return ( (const MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_WordCharacters();
}

void QsciLexerCoffeeScript_override_virtual_SetAutoIndentStyle(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__SetAutoIndentStyle = slot;
}

void QsciLexerCoffeeScript_virtualbase_SetAutoIndentStyle(void* self, int autoindentstyle) {
	( (MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_SetAutoIndentStyle(autoindentstyle);
}

void QsciLexerCoffeeScript_override_virtual_SetColor(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__SetColor = slot;
}

void QsciLexerCoffeeScript_virtualbase_SetColor(void* self, QColor* c, int style) {
	( (MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_SetColor(c, style);
}

void QsciLexerCoffeeScript_override_virtual_SetEolFill(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__SetEolFill = slot;
}

void QsciLexerCoffeeScript_virtualbase_SetEolFill(void* self, bool eoffill, int style) {
	( (MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_SetEolFill(eoffill, style);
}

void QsciLexerCoffeeScript_override_virtual_SetFont(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__SetFont = slot;
}

void QsciLexerCoffeeScript_virtualbase_SetFont(void* self, QFont* f, int style) {
	( (MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_SetFont(f, style);
}

void QsciLexerCoffeeScript_override_virtual_SetPaper(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__SetPaper = slot;
}

void QsciLexerCoffeeScript_virtualbase_SetPaper(void* self, QColor* c, int style) {
	( (MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_SetPaper(c, style);
}

void QsciLexerCoffeeScript_override_virtual_ReadProperties(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__ReadProperties = slot;
}

bool QsciLexerCoffeeScript_virtualbase_ReadProperties(void* self, QSettings* qs, struct miqt_string prefix) {
	return ( (MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_ReadProperties(qs, prefix);
}

void QsciLexerCoffeeScript_override_virtual_WriteProperties(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( (QsciLexerCoffeeScript*)(self) )->handle__WriteProperties = slot;
}

bool QsciLexerCoffeeScript_virtualbase_WriteProperties(const void* self, QSettings* qs, struct miqt_string prefix) {
	return ( (const MiqtVirtualQsciLexerCoffeeScript*)(self) )->virtualbase_WriteProperties(qs, prefix);
}

void QsciLexerCoffeeScript_Delete(QsciLexerCoffeeScript* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQsciLexerCoffeeScript*>( self );
	} else {
		delete self;
	}
}
