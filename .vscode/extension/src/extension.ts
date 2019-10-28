'use strict';
import * as vscode from 'vscode';
import * as path from 'path';
import * as cp from 'child_process';

export function activate(context: vscode.ExtensionContext) {
	const collection = vscode.languages.createDiagnosticCollection();

	if (vscode.window.activeTextEditor) {
		collection.clear()
		updateDiagnostics(vscode.window.activeTextEditor.document, collection);
	}
	context.subscriptions.push(vscode.window.onDidChangeActiveTextEditor(editor => {
		if (editor) {
			collection.clear()
			updateDiagnostics(editor.document, collection);
		}
	}));
}

function updateDiagnostics(document: vscode.TextDocument, collection: vscode.DiagnosticCollection): void {
	if (document && path.basename(document.uri.fsPath) === 'step.yml') {
		cp.exec('step-yml-linter -file '+document.uri.fsPath, (err, stdout, stderr) => {
			var lines = stdout.split('\n');
			for(var i = 0;i < lines.length;i++){
				if(lines[i]) {
					var components = lines[i].split(':');
			console.log(components)
					var file = components[0];
					var lineNo = Math.max(0, parseInt(components[1])-1);
					var column = Math.max(0, parseInt(components[2])-1);
					var string = components[3];

					collection.set(document.uri, [{
						code: '',
						message: string,
						range: new vscode.Range(new vscode.Position(lineNo, column), new vscode.Position(lineNo, 5000)),
						severity: vscode.DiagnosticSeverity.Error,
						source: '',
						// relatedInformation: [
						// 	new vscode.DiagnosticRelatedInformation(new vscode.Location(document.uri, new vscode.Range(new vscode.Position(1, 8), new vscode.Position(1, 9))), 'first assignment to `x`')
						// ]
					}]);
					collection = vscode.languages.createDiagnosticCollection();
				}
			}
			if (err) {
				console.log('error: ' + err);
			}
		});
	} else {
		collection.clear();
	}
}

// this method is called when your extension is deactivated
export function deactivate() {
}